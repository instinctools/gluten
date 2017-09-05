package gorm

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/jinzhu/gorm"
)

//TODO - fulfill all methods
type GormExecutionsRepo struct {
	connection *gorm.DB
}

func NewGormExecutionsRepo(URL string) *GormExecutionsRepo {
	return &GormExecutionsRepo{
		InitDb(URL),
	}
}

func (repo *GormExecutionsRepo) Create(execution *core.Execution) {
	tx := repo.connection.Begin()
	err := tx.Create(NewExecution(execution)).Error
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error during serving")
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func (repo *GormExecutionsRepo) Get(limit int, offset int) []core.Execution {
	return nil
}

func (repo *GormExecutionsRepo) GetById(id string) core.Execution {
	return core.Execution{}
}

func (repo *GormExecutionsRepo) Update(result core.Execution) {

}
