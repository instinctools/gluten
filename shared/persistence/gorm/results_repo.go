package gorm

import (
	"bitbucket.org/instinctools/gluten/core"
	//	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/jinzhu/gorm"
)

//TODO - fulfill all methods

type GormResultsRepo struct {
	connection *gorm.DB
}

func (repo *GormResultsRepo) Create(result core.StepResult) {
	tx := repo.connection.Begin()
	err := tx.Create(NewExecutionResult(result)).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
}

func (repo *GormResultsRepo) Get(limit int, offset int) []core.StepResult {
	return nil
}

func (repo *GormResultsRepo) GetById(id string) core.StepResult {
	return core.StepResult{}
}

func (repo *GormResultsRepo) Update(result core.StepResult) {

}
