package core

import "github.com/google/uuid"

type DefaultRunner struct {
	Handler ResultHandler
}

func (runner *DefaultRunner) Run(runnable Runnable) {
	executionID := uuid.New().String()
	stepResults := runnable.Run()
	for _, stepResult := range stepResults {
		stepResult.ExecutionID = executionID
		runner.Handler.Handle(stepResult)
	}
}
