package service

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/persistence"
	//	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/google/uuid"
)

type ExecutionService struct {
	runner        core.TestRunner
	executionRepo persistence.ExecutionRepo
}

func NewExecutionService(runner core.TestRunner, executionRepo persistence.ExecutionRepo) *ExecutionService {
	return &ExecutionService{
		runner,
		executionRepo,
	}
}

func (service *ExecutionService) ExecuteProject(project *core.Project) {
	execution := &core.Execution{
		ID:     uuid.New().String(),
		Status: core.STATUS_CREATED,
	}
	service.executionRepo.Create(execution)
	service.runner.Run(execution, project)
}
