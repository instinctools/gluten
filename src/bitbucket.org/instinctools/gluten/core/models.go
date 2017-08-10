package core

import (
)

//Common
type Common struct {
	Name string
}

// Project
type Project struct {
	Common
	Scenarios []TestScenario
}

func (p *Project) Add(ts TestScenario) {
	p.Scenarios = append(p.Scenarios, ts)
}

// Scenario
type TestScenario struct {
	Common
	Cases []TestCase
}

func (ts *TestScenario) Add(tc TestCase) {
	ts.Cases = append(ts.Cases, tc)
}

// Case
type TestCase struct {
	Common
	Steps []TestStep
}

func (tcase *TestCase) Add(step TestStep) {
	tcase.Steps = append(tcase.Steps, step)
}

// Step
type RunStep func(TestStep) (string, []Metric)
type BeforeStep func(TestStep)

func DoNothing(step TestStep){
}

type TestStep struct {
	Common
	BeforeF BeforeStep
	RunF RunStep
	Parameters map[string]string
	Substeps []TestStep
}

// Result
type StepResult struct {
	RunID       string
	Status      string
	ElapsedTime int16
	StepType    string
	Metrics     []Metric
}

// Metric
type Metric struct {
	Key string
	Val string
}

// TestRunner
type TestRunner interface {
	Run(c interface{})
}

// ResultHandler
type ResultHandler interface {
	Handle(result StepResult)
}

