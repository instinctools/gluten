package steps

import "bitbucket.org/instinctools/gluten/core"

//RepeaterStep ...
type RepeaterStep struct {
	CompositeStep
	repeats int
}

func NewRepeaterStep(name string, repeatsCount int) *RepeaterStep {
	return &RepeaterStep{
		CompositeStep{core.BaseStep{Name: name}},
		repeatsCount,
	}
}

func (step *RepeaterStep) Run() []core.StepResult {
	stepResults := []core.StepResult{}
	for i := 0; i < step.repeats; i++ {
		stepResults = append(stepResults, step.Run()...)
	}
	return stepResults
}
