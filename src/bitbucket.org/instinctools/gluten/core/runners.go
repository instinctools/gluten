package core

import (

)

type DefaultRunner struct {
	Handler ResultHandler
}

func (runner *DefaultRunner) Run(c interface{}) {
	//TODO - fix code dup in switch
	switch c.(type) {
	case Project:
		for _, element := range c.(Project).Scenarios {
			runner.Run(element)
		}
	case TestScenario:
		for _, element := range c.(TestScenario).Cases {
			runner.Run(element)
		}
	case TestCase:
		for _, element := range c.(TestCase).Steps {
			runner.Run(element)
		}
	case TestStep:
		step := c.(TestStep)
		stepType, metrics := step.RunF(step)
		runner.Handler.Handle(StepResult{
				Metrics: metrics,
				ElapsedTime: 1,
				RunID: "id1",
				Status: "Completed",
				StepType: stepType,
		})
	default:
		panic("Unknow type for running")
	}
}

