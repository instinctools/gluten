package steps


//CompositeStep ...
type CompositeStep struct {
	BaseStep
}

func NewCompositeStep(name string, subSteps []Step) *CompositeStep {
	return &CompositeStep{
		BaseStep{Name: name, SubSteps: subSteps},
	}
}

func (step *CompositeStep) Run() []StepResult {
	stepResults := []StepResult{}
	for _, s := range step.GetSubSteps() {
		stepResults = append(stepResults, s.Run()...)
	}
	return stepResults
}

func (step *CompositeStep) BeforeStep() {
	//validate and preset parameters
}
