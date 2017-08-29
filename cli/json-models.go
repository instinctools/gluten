package main

type Project struct {
	Name      string      `json:"Name"`
	Scenarios []TestScenario `json:"Scenarios"`
}

type TestScenario struct {
	Name      string     `json:"Name"`
	Cases []TestCase `json:"Cases"`
}

type TestCase struct {
	Name      string     `json:"Name"`
	Steps []TestStep `json:"Steps"`
}

type TestStep struct {
	Type   string   `json:"type"`
	Parameters map[string]string `json:"Parameters"`
}