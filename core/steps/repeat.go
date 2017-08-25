package steps

import (
	"bitbucket.org/instinctools/gluten/core"
)
//RepeaterStep ...
type RepeaterStep struct {
	CompositeStep
	repeats int
}

func NewRepeaterStep(name string, repeatsCount int, subSteps []core.Step) *RepeaterStep {
	return &RepeaterStep{
		CompositeStep{
			core.BaseStep{name},
			subSteps,
		},
		repeatsCount,
	}
}

// TODO fix. infinite loop
func (step *RepeaterStep) Run() []core.StepResult {
	stepResults := []core.StepResult{}
	for i := 0; i < step.repeats; i++ {
		stepResults = append(stepResults, step.Run()...)
	}
	return stepResults
}
