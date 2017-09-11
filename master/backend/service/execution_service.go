package service

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/persistence"
	//	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/google/uuid"
	"bitbucket.org/instinctools/gluten/core/result_handlers"
	"bitbucket.org/instinctools/gluten/shared/persistence/gorm"
)

type ExecutionService struct {
	runner        core.TestRunner
	executionRepo persistence.ExecutionRepo
}

var (
	ExecutionServiceInstance *ExecutionService
)

func init() {
	runner := core.NewDefaultRunner(&result_handlers.LoggableResultHandler{})
	ExecutionServiceInstance = &ExecutionService{
		runner,
		gorm.ExecutionsRepoInstance,
	}
}

func (service *ExecutionService) ExecuteProject(project *core.Project) *core.Execution {
	execution := &core.Execution{
		ID:     uuid.New().String(),
		Status: core.STATUS_CREATED,
	}
	service.executionRepo.Create(execution)
	service.runner.Run(execution, project)
	return execution
}
