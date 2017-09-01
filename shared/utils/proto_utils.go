package utils

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/steps"
	pb "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	pm "bitbucket.org/instinctools/gluten/shared/rpc/master"
	"encoding/json"
	"strconv"
	"fmt"
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
				step := ParseProto2Step(pStep)
				testCase.Add(step)
			}
			testScenario.Add(testCase)
		}
		testProject.Add(testScenario)
	}
	return &testProject
}

func ParseStep2Proto(step *core.TestStep) *pm.Step {
	var pSubSteps []pm.TestStep
	for _, subStep := range step.GetSubSteps() {
		pSubStep := parseProtoStep(subStep)
		pSubSteps = append(pSubSteps, pSubStep)
	}
	sMap := parsIMap(step.GetParams())
	pStep := pm.Step{Name:step.GetCommon().Name, Type:step.GetStepType(), Parameters:sMap, pSubSteps}
	return pStep
}

func ParseProto2Step(pStep *pb.Step) core.TestStep {
	var subSteps []core.TestStep
	for _, pSubStep := range pStep.GetSubSteps() {
		subStep := parseProtoStep(pSubStep)
		subSteps = append(subSteps, subStep)
	}
	iMap := parsPMap(pStep.Parameters)
	step := steps.NewStep(pStep.Type, pStep.Name, iMap, subSteps)
	return step
}

func ParseProto2Step(pStep *pm.Step) core.TestStep {
	var subSteps []core.TestStep
	for _, pSubStep := range pStep.GetSubSteps() {
		subStep := parseProtoStep(pSubStep)
		subSteps = append(subSteps, subStep)
	}
	iMap := parsPMap(pStep.Parameters)
	step := steps.NewStep(pStep.Type, pStep.Name, iMap, subSteps)
	return step
}

func parsIMap(iMap map[string]interface) map[string]string {
	sMap := make(map[string]string)
	for k, v := range iMap {
		sMap[k] = string(v)
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
