package steps

//RepeaterStep ...
type RepeaterStep struct {
	CompositeStep
	repeats int
}

func NewRepeaterStep(name string, repeatsCount int, subSteps []Step) *RepeaterStep {
	return &RepeaterStep{
		CompositeStep{
			BaseStep{name},
			subSteps,
		},
		repeatsCount,
	}
}

// TODO fix. infinite loop
func (step *RepeaterStep) Run() []StepResult {
	stepResults := []StepResult{}
	for i := 0; i < step.repeats; i++ {
		stepResults = append(stepResults, step.Run()...)
	}
	return stepResults
}
