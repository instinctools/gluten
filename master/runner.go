package main

import (
	"time"

	"bitbucket.org/instinctools/gluten/core"
)

func RunTestArchitecture() {
	params1 := make(map[string]interface{})
	params1["URL"] = "https://google.com"

	params2 := make(map[string]interface{})
	params2["REPEATS"] = 5

	params3 := make(map[string]interface{})
	params3["PARALLEL_REQUESTS"] = 4

	case1 := core.TestCase{
		Common: core.Common{Name: "Case1"},
		Steps: []core.TestStep{
			core.NewStep(core.GetRequestStepConstant, "Step1", params1, []core.TestStep{}),
		}}

	scenario1 := core.TestScenario{Common: core.Common{Name: "Sc1"}}
	scenario1.Add(case1)

	project1 := core.Project{Common: core.Common{Name: "Sc1"}}
	project1.Add(scenario1)

	runner := core.DefaultRunner{
		Handler: core.LoggableResultHandler{Name: "LoggableHandler1"},
	}
	runner.Run(project1)

	time.Sleep(5 * time.Second)

}
