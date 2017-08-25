package core

import (
	"reflect"
)

type Runnable interface {
	Run() []StepResult
}

// Project ...
type Project struct {
	Name      string
	Scenarios []Scenario
}

func (p *Project) Run() (stepResults []StepResult) {
	for _, scenario := range p.Scenarios {
		stepResults = append(stepResults, scenario.Run()...)
	}
	return
}

func (p *Project) Add(ts Scenario) {
	p.Scenarios = append(p.Scenarios, ts)
}

func (p *Project) GetSteps() []Step {
	steps := []Step{}
	for _, scenario := range p.Scenarios {
		for _, testCase := range scenario.Cases {
			for _, step := range testCase.Steps {
				steps = append(steps, collectSubSteps(step)...)
			}
		}
	}
	return steps
}

func collectSubSteps(t Step) []Step {
	_, ok := reflect.TypeOf(t).MethodByName("GetSubSteps")
	if ok {
		for _, stepValue := range reflect.ValueOf(t).MethodByName("GetSubSteps").Call([]reflect.Value{}) {
			steps := stepValue.Interface().([]Step)
			for _, step := range steps {
				if len(reflect.ValueOf(step).MethodByName("GetSubSteps").Call([]reflect.Value{})) != 0 {
					return append(collectSubSteps(step), step)
				}
			}
		}
	}
	return []Step{t}
}

// Scenario ...
type Scenario struct {
	Name  string
	Cases []TestCase
}

func (s *Scenario) Run() (stepResults []StepResult) {
	for _, testCase := range s.Cases {
		stepResults = append(stepResults, testCase.Run()...)
	}
	return
}

func (s *Scenario) Add(tc TestCase) {
	s.Cases = append(s.Cases, tc)
}

// TestCase ...
type TestCase struct {
	Name  string
	Steps []Step
}

func (c *TestCase) Run() (stepResults []StepResult) {
	for _, step := range c.Steps {
		stepResults = append(stepResults, step.Run()...)
	}
	return
}

func (c *TestCase) Add(step Step) {
	c.Steps = append(c.Steps, step)
}

// Step ...
type Step interface {
	Runnable
	GetName() string
	BeforeStep()
}

type BaseStep struct {
	Name     string
}

func (step *BaseStep) GetName() string {
	return step.Name
}

func (step *BaseStep) BeforeStep() {
	//Do nothing
}
 
// StepResult ...
type StepResult struct {
	ExecutionID string
	//TODO make me as enum
	Status  string
	Error   error
	Step    Step
	Metrics []Metric
}

// Metric ...
type Metric struct {
	Key string
	Val interface{}
}

// ResultHandler ...
type ResultHandler interface {
	Handle(result StepResult)
}
