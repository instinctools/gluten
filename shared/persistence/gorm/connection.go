package gorm

import (
	"bitbucket.org/instinctools/gluten/master/backend/config"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Connection struct {
	gorm *gorm.DB
}

var (
	connectionFactory *Connection
)

func init() {
	connectionFactory = &Connection{
		gorm: InitDb(),
	}
}

func InitDb() *gorm.DB {
	db, err := gorm.Open("postgres", config.GlobalConfig.DB.Connection.URL)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("gorm.Open error")
	}
	db.LogMode(true)
	return db
}
