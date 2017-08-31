package client

type Project struct {
	Name      string         `json:"Name"`
	Scenarios []TestScenario `json:"Scenarios"`
}

type TestScenario struct {
	Name  string     `json:"Name"`
	Cases []TestCase `json:"Cases"`
}

type TestCase struct {
	Name  string     `json:"Name"`
	Steps []TestStep `json:"Steps"`
}

type TestStep struct {
	Name       string            `json:"Name"`
	Type       string            `json:"Type"`
	Parameters map[string]string `json:"Parameters"`
	SubSteps   []TestStep        `json:"SubSteps"`
}
