package gorm

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/master/backend/config"
	"github.com/jinzhu/gorm"
)

type ResultsRepo struct {
	connection *gorm.DB
}

var (
	GetResultsRepo *ResultsRepo
)

func init() {
	GetResultsRepo = &ResultsRepo{
		InitDb(config.GlobalConfig.DB.Connection.URL),
	}
}
func (repo *ResultsRepo) Create(result core.StepResult) {
	tx := repo.connection.Begin()
	err := tx.Create(NewExecutionResult(result)).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
}

func (repo *ResultsRepo) Get(limit int, offset int) []core.StepResult {
	var dto []Result
	repo.connection.
		Limit(limit).
		Offset(offset).
		Find(&dto)
	results := []core.StepResult{}
	for _, elem := range dto {
		results = append(results, *elem.toStepResult())
	}
	return results
}

func (repo *ResultsRepo) GetByExecutionId(id string, limit int, offset int) []core.StepResult {
	var dto []Result
	repo.connection.
		Preload("Metrics").
		Limit(limit).
		Offset(offset).
		Where("execution_id = ?", id).
		Find(&dto)
	results := []core.StepResult{}
	for _, elem := range dto {
		results = append(results, *elem.toStepResult())
	}
	return results
}

func (repo *ResultsRepo) GetById(id string) core.StepResult {
	var dto Result
	repo.connection.First(&dto, id)
	return *dto.toStepResult()
}

func (repo *ResultsRepo) Update(result core.StepResult) {
	repo.connection.Find(&Result{}).Update(result)
}
