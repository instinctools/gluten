package gorm

import (
	"bitbucket.org/instinctools/gluten/core"
	"github.com/jinzhu/gorm"
)

type ResultsRepo struct {
	connection *gorm.DB
	rawRepo    *RawResultsRepo
}

type RawResultsRepo struct {
	connection *gorm.DB
}

var (
	ResultsRepoInstance *ResultsRepo
	RawResultsRepoInstance *RawResultsRepo
)

func init() {
	RawResultsRepoInstance = &RawResultsRepo{
		connection: connectionFactory.gorm,
	}

	ResultsRepoInstance = &ResultsRepo{
		connection: connectionFactory.gorm,
		rawRepo: RawResultsRepoInstance,
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
	dto = repo.rawRepo.Get(limit, offset)
	results := []core.StepResult{}
	for _, elem := range dto {
		results = append(results, *elem.toStepResult())
	}
	return results
}

func (repo *RawResultsRepo) Get(limit int, offset int) []Result {
	var dto []Result
	repo.connection.
		Limit(limit).
		Offset(offset).
		Find(&dto)
	return dto
}

func (repo *ResultsRepo) GetByExecutionId(id string, limit int, offset int) []core.StepResult {
	var dto []Result
	dto = repo.rawRepo.GetByExecutionId(id, limit, offset)
	results := []core.StepResult{}
	for _, elem := range dto {
		results = append(results, *elem.toStepResult())
	}
	return results
}

func (repo *RawResultsRepo) GetByExecutionId(id string, limit int, offset int) []Result {
	var dto []Result
	repo.connection.
		Preload("Metrics").
		Limit(limit).
		Offset(offset).
		Where("execution_id = ?", id).
		Find(&dto)
	return dto
}

func (repo *ResultsRepo) GetById(id string) core.StepResult {
	var dto Result
	repo.connection.First(&dto, id)
	return *dto.toStepResult()
}

func (repo *ResultsRepo) Update(result core.StepResult) {
	repo.connection.Find(&Result{}).Update(result)
}
