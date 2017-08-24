package steps

import "bitbucket.org/instinctools/gluten/core"

//CompositeStep ...
type CompositeStep struct {
	core.BaseMultipleStep
}

func NewCompositeStep(name string, subSteps []core.Step) *CompositeStep {
	return &CompositeStep{
		core.BaseMultipleStep{
			core.BaseStep{name},
			subSteps,
		},
	}
}

func (step *CompositeStep) Run() []core.StepResult {
	stepResults := []core.StepResult{}
	for _, s := range step.GetSubSteps() {
		stepResults = append(stepResults, s.Run()...)
	}
	return stepResults
}

func (step *CompositeStep) BeforeStep() {
	//validate and preset parameters
}
