package prst

import (
	"github.com/jinzhu/gorm"
	"github.com/mattes/migrate"
	dStub "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"log"
	"path/filepath"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=1 dbname=gluten_db sslmode=disable")
	CheckErr(err, "gorm.Open failed")
	//	db.LogMode(true)
	driver, err := dStub.WithInstance(db.DB(), &dStub.Config{})
	path, err := filepath.Abs("./share/migrations/")
	m, err := migrate.NewWithDatabaseInstance("file://"+path, "postgres", driver)

	m.Steps(1)
	return db
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
