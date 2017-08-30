package utils

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/steps"
	pb "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	"encoding/json"
	"strconv"
)

func DeserializeJsonToProto(jsonProject string) *pb.Project {
	deserializedProject := pb.Project{}
	err := json.Unmarshal([]byte(jsonProject), &deserializedProject)
	if err != nil {
		panic(err)
	}
	return &deserializedProject
}

func ParseProto2Project(pProject *pb.Project) *core.Project {
	testProject := core.Project{Common: core.Common{pProject.Name}}
	for _, pScenario := range pProject.GetScenarios() {
		testScenario := core.TestScenario{Common: core.Common{pScenario.Name}}
		for _, pCase := range pScenario.GetCases() {
			testCase := core.TestCase{Common: core.Common{pCase.Name}}
			for _, pStep := range pCase.GetSteps() {
				step := parseProtoStep(pStep)
				testCase.Add(step)
			}
			testScenario.Add(testCase)
		}
		testProject.Add(testScenario)
	}
	return &testProject
}

func parseProtoStep(pStep *pb.Step) core.TestStep {
	var subSteps []core.TestStep
	for _, pSubStep := range pStep.GetSubSteps() {
		subStep := parseProtoStep(pSubStep)
		subSteps = append(subSteps, subStep)
	}
	iMap := parsMap(pStep.Parameters)
	step := steps.NewStep(pStep.Type, pStep.Name, iMap, subSteps)
	return step
}

func parsMap(oMap map[string]string) map[string]interface{} {
	iMap := make(map[string]interface{})
	for k, v := range oMap {
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
