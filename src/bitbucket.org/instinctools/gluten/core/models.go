//package core

package main

import (
	log "github.com/sirupsen/logrus"
	"os"
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

func (p *Project) add(ts TestScenario) {
	p.Scenarios = append(p.Scenarios, ts)
}

// Scenario
type TestScenario struct {
	Common
	Cases []TestCase
}

func (ts *TestScenario) add(tc TestCase) {
	ts.Cases = append(ts.Cases, tc)
}

// Case
type TestCase struct {
	Common
	Steps []TestStep
}

func (tcase *TestCase) add(step TestStep) {
	log.WithFields(log.Fields{
		"Case": tcase.Name,
		"Step": step.Name,
	}).Debug("Adding step to case")
	tcase.Steps = append(tcase.Steps, step)
}

// Step
type StepTyper interface {
	getType() string
}

type StepProtoType struct {
	Type string
}

func (p StepProtoType) getType() string {
	return p.Type
}

type runStep func(TestStep) (StepTyper, []Metric)

type TestStep struct {
	Common
	runF runStep
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
	key string
	val string
}

// TestRunner
type TestRunner interface {
	Run(c interface{})
}

type DefaultRunner struct {
	handler ResultHandler
}

func (runner *DefaultRunner) Run(c interface{}) {
	//TODO - fix code dup in switch
	switch c.(type) {
	case Project:
		for _, element := range c.(Project).Scenarios {
			runner.Run(element)
		}
	case TestScenario:
		for _, element := range c.(TestScenario).Cases {
			runner.Run(element)
		}
	case TestCase:
		for _, element := range c.(TestCase).Steps {
			runner.Run(element)
		}
	case TestStep:
		step := c.(TestStep)
		typer, metrics := step.runF(step)
		runner.handler.Handle(StepResult{
				Metrics: metrics,
				ElapsedTime: 1,
				RunID: "id1",
				Status: "Completed",
				StepType: typer.getType(),
		})
	default:
		panic("Unknow type for running")
	}
}

// ResultHandler
type ResultHandler interface {
	Handle(result StepResult)
}

type LoggableResultHandler struct {
	Name string
}

func (h LoggableResultHandler) Handle(result StepResult) {
	log.WithFields(log.Fields{
		"result": result,
	}).Info("Step has been handled")
}

// Main - TODO: removed from here
func EchoStep(t TestStep) (StepTyper, []Metric) {
	log.WithFields(log.Fields{
		"step": t,
	}).Info("Step has been invoked")
	typer := struct {StepProtoType}{StepProtoType{Type:"ECHO_STEP"}} 
	return typer, []Metric{}
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	case1 := TestCase{
		Common: Common{Name: "Case1"},
		Steps:     []TestStep{TestStep{Common: Common{Name: "Step1"}, runF: EchoStep}},
	}
	scenario1 := TestScenario{Common: Common{Name: "Sc1"}}
	scenario1.add(case1)

	project1 := Project{Common: Common{Name: "Sc1"}}
	project1.add(scenario1)

	runner := DefaultRunner{
		handler: LoggableResultHandler{Name: "LoggableHandler1"},
	}
	runner.Run(project1)
}
