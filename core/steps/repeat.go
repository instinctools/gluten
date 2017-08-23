package steps


//RepeaterStep ...
type RepeaterStep struct {
	CompositeStep
	repeats int
}

func NewRepeaterStep(name string, repeatsCount int) *RepeaterStep {
	return &RepeaterStep{
		CompositeStep{BaseStep{Name: name}},
		repeatsCount,
	}
}

func (step *RepeaterStep) Run() []StepResult {
	stepResults := []StepResult{}
	for i := 0; i < step.repeats; i++ {
		stepResults = append(stepResults, step.Run()...)
	}
	return stepResults
}
