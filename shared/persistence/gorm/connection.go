package gorm

import (
	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDb(URL string) *gorm.DB {
	db, err := gorm.Open("postgres", URL)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("gorm.Open error")
	}
	db.LogMode(true)
	return db
}