package core

import ()



type Project struct {
	Name      string
	Scenarios []TestScenario
}

func (p *Project) addScenario(ts TestScenario) {
	p.Scenarios = append(p.Scenarios, ts)
}

type TestScenario struct {
	Name  string
	Cases []TestCase
}

type TestCase struct {
	Name  string
	Steps []TestStep
}

type TestStep struct {
	Name string
}
