package utils

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/steps"
	rpcCli "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	rpcMaster "bitbucket.org/instinctools/gluten/shared/rpc/master"
	"encoding/json"
	"strconv"
)

func DeserializeJsonToProto(jsonProject string) *rpcCli.Project {
	deserializedProject := rpcCli.Project{}
	err := json.Unmarshal([]byte(jsonProject), &deserializedProject)
	if err != nil {
		panic(err)
	}
	return &deserializedProject
}

func ParseProto2Project(pProject *rpcCli.Project) *core.Project {
	testProject := &core.Project{Common: core.Common{pProject.Name}}
	for _, pScenario := range pProject.GetScenarios() {
		testScenario := core.TestScenario{Common: core.Common{pScenario.Name}}
		for _, pCase := range pScenario.GetCases() {
			testCase := core.TestCase{Common: core.Common{pCase.Name}}
			for _, pStep := range pCase.GetSteps() {
				step := ParseCliProto2Step(pStep)
				testCase.Add(step)
			}
			testScenario.Add(testCase)
		}
		testProject.Add(testScenario)
	}
	return testProject
}

func ParseStep2Proto(context *core.Execution, step core.TestStep) *rpcMaster.Step {
	var pSubSteps []*rpcMaster.Step
	for _, subStep := range step.GetSubSteps() {
		pSubStep := ParseStep2Proto(context, subStep)
		pSubSteps = append(pSubSteps, pSubStep)
	}
	sMap := parsIMap(step.GetParams())
	return &rpcMaster.Step{
		Name:       step.GetCommon().Name,
		Type:       step.GetStepType(),
		Parameters: sMap,
		SubSteps:   pSubSteps,
		Exec:       &rpcMaster.Execution{ID: context.ID, Status: context.Status}}
}

func ParseCliProto2Step(pStep *rpcCli.Step) core.TestStep {
	var subSteps []core.TestStep
	for _, pSubStep := range pStep.GetSubSteps() {
		subStep := ParseCliProto2Step(pSubStep)
		subSteps = append(subSteps, subStep)
	}
	iMap := parsPMap(pStep.Parameters)
	step := steps.NewStep(pStep.Type, pStep.Name, iMap, subSteps)
	return step
}

func ParseMasterProto2Step(pStep *rpcMaster.Step) (*core.Execution, core.TestStep) {
	var subSteps []core.TestStep
	for _, pSubStep := range pStep.GetSubSteps() {
		_, subStep := ParseMasterProto2Step(pSubStep)
		subSteps = append(subSteps, subStep)
	}
	iMap := parsPMap(pStep.Parameters)
	step := steps.NewStep(pStep.Type, pStep.Name, iMap, subSteps)
	return &core.Execution{ID: pStep.Exec.ID, Status: pStep.Exec.Status}, step
}

func parsIMap(iMap map[string]interface{}) map[string]string {
	sMap := make(map[string]string)
	for k, v := range iMap {
		sMap[k] = v.(string)
	}
	return sMap
}

func parsPMap(sMap map[string]string) map[string]interface{} {
	iMap := make(map[string]interface{})
	for k, v := range sMap {
		iMap[k] = parsString(v)
	}
	return iMap
}

func parsString(str string) interface{} {
	b, err := strconv.ParseBool(str)
	if err == nil {
		return b
	}
	f, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return f
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return i
	}
	u, err := strconv.ParseUint(str, 10, 64)
	if err == nil {
		return u
	}
	return str
}
