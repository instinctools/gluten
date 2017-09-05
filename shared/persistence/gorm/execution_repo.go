package gorm

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/jinzhu/gorm"
)

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
	var dto []Execution
	repo.connection.
		Limit(limit).
		Offset(offset).
		Find(&dto)
	executions := []core.Execution{}
	for _, elem := range dto {
		executions = append(executions, *elem.toExecution())
	}
	return executions
}

func (repo *GormExecutionsRepo) GetById(id string) core.Execution {
	var dto Execution
	repo.connection.First(&dto, id)
	return *DtoToExecution(&dto)
}

func (repo *GormExecutionsRepo) Update(execution core.Execution) {
	repo.connection.Find(&Execution{}).Update(execution)
}

func (repo *GormExecutionsRepo) Delete(execution core.Execution) {
	repo.connection.Find(&Execution{}).Delete(execution)
}
