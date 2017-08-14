package core

import (

)

type DefaultRunner struct {
	Handler ResultHandler
}

func (runner *DefaultRunner) Run(c interface{}) {
	//TODO - generate id here
}

func (runner *DefaultRunner) run1(c interface{}, execId string){
	//TODO - fix code dup in switch
	switch c.(type) {
	case Project:
		for _, element := range c.(Project).Scenarios {
			runner.run1(element, execId)
		}
	case TestScenario:
		for _, element := range c.(TestScenario).Cases {
			runner.run1(element, execId)
		}
	case TestCase:
		for _, element := range c.(TestCase).Steps {
			runner.run1(element, execId)
		}
	case TestStep:
		step := c.(TestStep)
		metrics := step.Run()
		runner.Handler.Handle(StepResult{
				Metrics: metrics,
				ExecutionID: execId,
				Status: "Completed",
				StepType: step.GetStepType(),
		})
	default:
		panic("Unknow type for running")
	}
	
}

