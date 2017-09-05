package core

import (
	"bitbucket.org/instinctools/gluten/shared/logging"
	"errors"
)

var (
	ErrUnknowType = errors.New("Unknow type to run")
)

type DefaultRunner struct {
	handler ResultHandler
}

func NewDefaultRunner(handler ResultHandler) TestRunner {
	return &DefaultRunner{handler}
}

func (runner *DefaultRunner) Run(context *Execution, tests interface{}) error {
	logging.WithFields(logging.Fields{
		"to_run": tests,
	}).Info("Trying to run tests")
	return runner.run1(context, tests)
}

func (runner *DefaultRunner) run1(context *Execution, c interface{}) error {
	//TODO - fix code dup in switch
	switch c.(type) {
	case *Project:
		p := c.(*Project)
		for _, element := range p.Scenarios {
			runner.run1(context, element)
		}
	case Project:
		p := c.(Project)
		for _, element := range p.Scenarios {
			runner.run1(context, element)
		}
	case TestScenario:
		for _, element := range c.(TestScenario).Cases {
			runner.run1(context, element)
		}
	case TestCase:
		for _, element := range c.(TestCase).Steps {
			runner.run1(context, element)
		}
	case TestStep:
		step := c.(TestStep)
		metrics := step.Run(context)
		runner.handler.Handle(StepResult{
			ExecutionID: context.ID,
			Metrics:     metrics,
			Status:      "Completed",
			StepType:    step.GetStepType(),
		})
	default:
		logging.WithFields(logging.Fields{
			"to_run": c,
		}).Error("Can't run struct")
		return ErrUnknowType
	}
	return nil
}
