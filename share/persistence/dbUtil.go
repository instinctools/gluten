package persistence

import (
	"github.com/jinzhu/gorm"
	"log"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=1 dbname=gluten_db sslmode=disable")
	CheckErr(err, "gorm.Open failed")
//	db.LogMode(true)

	db.DropTable(Execution{})
	db.DropTable(Metric{})
	db.DropTable(ExecutionResult{})

	db.AutoMigrate(&ExecutionResult{}, &Execution{}, &Metric{})

	db.Model(&Execution{}).AddForeignKey("result_id", "executions_results(id)", "CASCADE", "CASCADE")
	db.Model(&Metric{}).AddForeignKey("execution_result_id", "executions_results(id)", "CASCADE", "CASCADE")

	return db
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
