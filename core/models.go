package core

// Common ...
type Common struct {
	Name string
}

// Project ...
type Project struct {
	Common
	Scenarios []TestScenario
}

func (p *Project) Add(ts TestScenario) {
	p.Scenarios = append(p.Scenarios, ts)
}

func (p *Project) GetAllSteps() []TestStep {
	steps := []TestStep{}
	for _, scenario := range p.Scenarios {
		for _, tcase := range scenario.Cases {
			for _, step := range tcase.Steps {
				steps = append(steps, step)
				collectSubSteps(step, steps)
			}
		}
	}
	return steps
}

func collectSubSteps(t TestStep, accum []TestStep) {
	for _, tstep := range t.GetSubSteps() {
		accum = append(accum, tstep)
		if len(tstep.GetSubSteps()) != 0 {
			collectSubSteps(tstep, accum)
		}
	}
}

// TestScenario ...
type TestScenario struct {
	Common
	Cases []TestCase
}

func (ts *TestScenario) Add(tc TestCase) {
	ts.Cases = append(ts.Cases, tc)
}

// TestCase ...
type TestCase struct {
	Common
	Steps []TestStep
}

func (tcase *TestCase) Add(step TestStep) {
	tcase.Steps = append(tcase.Steps, step)
}

// TestStep ...
type TestStep interface {
	GetCommon() Common
	GetParams() map[string]interface{}
	GetSubSteps() []TestStep
	GetStepType() string

	BeforeStep()
	Run() []Metric
}

type BaseTestStep struct {
	Common
	Parameters map[string]interface{}
	Substeps   []TestStep
}

// StepResult ...
type StepResult struct {
	ExecutionID string
	Status      string
	StepType    string
	Metrics     []Metric
}

// Metric ...
type Metric struct {
	Key string
	Val interface{}
}

// TestRunner ...
type TestRunner interface {
	Run(c interface{})
}

// ResultHandler ...
type ResultHandler interface {
	Handle(result StepResult)
}
