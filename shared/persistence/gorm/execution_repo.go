package gorm

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/jinzhu/gorm"
)

type ExecutionRepoWrapper struct {
	rawRepo *RawExecutionRepo
}

func NewExecutionRepoWrapper(rawRepo *RawExecutionRepo) *ExecutionRepoWrapper {
	return &ExecutionRepoWrapper{
		rawRepo,
	}
}

func (repo *ExecutionRepoWrapper) Create(execution *core.Execution) {
	repo.rawRepo.Create(NewExecution(execution))
}

func (repo *ExecutionRepoWrapper) Get(limit int, offset int) []core.Execution {
	var dto []Execution
	dto = repo.rawRepo.Get(limit, offset)
	executions := []core.Execution{}
	for _, elem := range dto {
		executions = append(executions, *elem.toExecution())
	}
	return executions
}

func (repo *ExecutionRepoWrapper) GetById(id string) *core.Execution {
	return repo.rawRepo.GetById(id).toExecution()
}

func (repo *ExecutionRepoWrapper) Update(execution *core.Execution) {
	repo.rawRepo.Update(NewExecution(execution))
}

func (repo *ExecutionRepoWrapper) Delete(execution *core.Execution) {
	repo.rawRepo.Delete(NewExecution(execution))
}

//Used by Web-UI/Rest services
type RawExecutionRepo struct {
	connection *gorm.DB
}

func NewRawExecutionRepo(conn *Connection) *RawExecutionRepo {
	return &RawExecutionRepo{
		connection: conn.gorm,
	}
}

func (repo *RawExecutionRepo) Create(execution *Execution) {
	tx := repo.connection.Begin()
	err := tx.Create(execution).Error
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error during serving")
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func (repo *RawExecutionRepo) Get(limit int, offset int) []Execution {
	var dto []Execution
	repo.connection.
		Limit(limit).
		Offset(offset).
		Find(&dto)
	return dto
}

func (repo *RawExecutionRepo) GetById(id string) *Execution {
	var dto *Execution
	repo.connection.First(dto, id)
	return dto
}

func (repo *RawExecutionRepo) Update(execution *Execution) {
	repo.connection.Update(execution)
}

func (repo *RawExecutionRepo) Delete(execution *Execution) {
	repo.connection.Delete(execution)
}
