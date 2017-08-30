package steps

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/shared/logging"
)

// Step registry & Step factory
type newStepF func(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep

var stepFactories = make(map[string]newStepF)

func RegisterStepFactory(name string, factory newStepF) {
	if factory == nil {
		panic("Factory does not exist.")
	}
	stepFactories[name] = factory
	logging.WithFields(logging.Fields{
		"step": name,
	}).Info("Step has been registered")
}

func NewStep(step string, name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
	return stepFactories[step](name, params, substeps)
}
