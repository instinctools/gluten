package persistence

import (
	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/jinzhu/gorm"
	"github.com/mattes/migrate"
	dStub "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"path/filepath"
	"os"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=1 dbname=gluten_db sslmode=disable")
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("gorm.Open error")
	}
	db.LogMode(true)
	driver, err := dStub.WithInstance(db.DB(), &dStub.Config{})
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("WithInstance error")
	}
	project_path, _ := os.Getwd()

	path, err := filepath.Abs(project_path+"/gluten/migrations/")
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Abs error")
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+path, "postgres", driver)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("NewWithDatabaseInstance error")
	}
	m.Steps(1)
	return db
}
