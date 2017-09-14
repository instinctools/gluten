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
	projectStore  ProjectStore
}

func NewExecutionService(runner core.TestRunner, executionRepo persistence.ExecutionRepo, projectStore ProjectStore) *ExecutionService {
	return &ExecutionService{
		runner,
		executionRepo,
		projectStore,
	}
}

func (service *ExecutionService) ExecuteProject(project *core.Project) *core.Execution {
	execution := &core.Execution{
		ID:     uuid.New().String(),
		Status: core.STATUS_CREATED,
	}
	service.executionRepo.Create(execution)
	service.projectStore.AddProject(project)
	service.runner.Run(execution, project)
	return execution
}
