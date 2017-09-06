package gorm

import (
	"bitbucket.org/instinctools/gluten/core"
	//	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/jinzhu/gorm"
)

type GormResultsRepo struct {
	connection *gorm.DB
}

func NewGormResultsRepo(URL string) *GormResultsRepo {
	return &GormResultsRepo{
		InitDb(URL),
	}
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

func (repo *GormResultsRepo) GetByExecutionId(id string, limit int, offset int) []core.StepResult {
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

func (repo *GormResultsRepo) GetById(id string) core.StepResult {
	var dto Result
	repo.connection.First(&dto, id)
	return *dto.toStepResult()
}

func (repo *GormResultsRepo) Update(result core.StepResult) {
	repo.connection.Find(&Result{}).Update(result)
}
