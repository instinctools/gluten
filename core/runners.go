package core

import (
	"github.com/google/uuid"
	"bitbucket.org/instinctools/gluten/core/steps"
)	

type DefaultRunner struct {
	Handler steps.ResultHandler
}

func (runner *DefaultRunner) Run(runnable steps.Runnable) {
	executionID := uuid.New().String()
	stepResults := runnable.Run()
	for _, stepResult := range stepResults {
		stepResult.ExecutionID = executionID
		runner.Handler.Handle(stepResult)
	}
}
