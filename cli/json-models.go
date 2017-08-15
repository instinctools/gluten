package main

type Project struct {
	Name      string      `json:"name"`
	TestSuits []TestSuite `json:"testSuits"`
}

type TestSuite struct {
	Name      string     `json:"name"`
	TestCases []TestCase `json:"testCases"`
}

type TestCase struct {
	Name      string     `json:"name"`
	TestSteps []TestStep `json:"testSteps"`
}

type TestStep struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Params []Params `json:"params"`
}

type Params struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type JSONModel struct {
	Project Project `json:"project"`
}
