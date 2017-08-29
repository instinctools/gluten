package persistence

import (
	"log"
	"path/filepath"

	"github.com/jinzhu/gorm"
	"github.com/mattes/migrate"
	dStub "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=1 dbname=gluten_db sslmode=disable")
	CheckErr(err, "gorm.Open failed")
	db.LogMode(true)
	driver, err := dStub.WithInstance(db.DB(), &dStub.Config{})
	CheckErr(err, "dStub.WithInstance failed")
	path, err := filepath.Abs("./gluten/shared/persistence/migrations/")
	CheckErr(err, "filepath.Abs failed")
	m, err := migrate.NewWithDatabaseInstance("file://"+path, "postgres", driver)
	CheckErr(err, "migrate.NewWithDatabaseInstancefailed")
	m.Steps(1)
	return db
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
