package core

import "github.com/google/uuid"

type DefaultRunner struct {
	Handler ResultHandler
}

func (runner *DefaultRunner) Run(c interface{}) {
	executionID := uuid.New().String()
	runner.run1(c, executionID)
}

func (runner *DefaultRunner) run1(c interface{}, execID string) {
	//TODO - fix code dup in switch
	switch c.(type) {
	case Project:
		for _, element := range c.(Project).Scenarios {
			runner.run1(element, execID)
		}
	case TestScenario:
		for _, element := range c.(TestScenario).Cases {
			runner.run1(element, execID)
		}
	case TestCase:
		for _, element := range c.(TestCase).Steps {
			runner.run1(element, execID)
		}
	case TestStep:
		step := c.(TestStep)
		metrics := step.Run()
		runner.Handler.Handle(StepResult{
			Metrics:     metrics,
			ExecutionID: execID,
			Status:      "Completed",
			StepType:    step.GetStepType(),
		})
	default:
		panic("Unknow type for running")
	}

}
