package gorm

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/jinzhu/gorm"
)

type ExecutionsRepo struct {
	rawRepo    *RawExecutionsRepo
	connection *gorm.DB
}

type RawExecutionsRepo struct {
	connection *gorm.DB
}

var (
	ExecutionsRepoInstance    *ExecutionsRepo
	RawExecutionsRepoInstance *RawExecutionsRepo
)

func init() {
	RawExecutionsRepoInstance = &RawExecutionsRepo{
		connection: connectionFactory.gorm,
	}
	ExecutionsRepoInstance = &ExecutionsRepo{
		rawRepo:    RawExecutionsRepoInstance,
		connection: connectionFactory.gorm,
	}
}

func (repo *ExecutionsRepo) Create(execution *core.Execution) {
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

func (repo *ExecutionsRepo) Get(limit int, offset int) []core.Execution {
	var dto []Execution
	dto = repo.rawRepo.Get(limit, offset)
	executions := []core.Execution{}
	for _, elem := range dto {
		executions = append(executions, *elem.toExecution())
	}
	return executions
}

func (repo *RawExecutionsRepo) Get(limit int, offset int) []Execution {
	var dto []Execution
	repo.connection.
		Limit(limit).
		Offset(offset).
		Find(&dto)
	return dto
}

func (repo *ExecutionsRepo) GetById(id string) core.Execution {
	var dto Execution
	repo.connection.First(&dto, id)
	return *dto.toExecution()
}

func (repo *ExecutionsRepo) Update(execution core.Execution) {
	repo.connection.Find(&Execution{}).Update(execution)
}

func (repo *ExecutionsRepo) Delete(execution core.Execution) {
	repo.connection.Find(&Execution{}).Delete(execution)
}
