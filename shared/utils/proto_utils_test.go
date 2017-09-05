package utils_test

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/steps"
	rpcCli "bitbucket.org/instinctools/gluten/shared/rpc/cli"
	"bitbucket.org/instinctools/gluten/shared/utils"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestProtoUtils(t *testing.T) {
	pProject := ProtoJsonDeserializerTest(t)
	Proto2ProjectTest(t, pProject)
	Step2ProtoTest(t)

}

func ProtoJsonDeserializerTest(t *testing.T) *rpcCli.Project {
	project := utils.DeserializeJsonToProto(getTestProjectJson())
	assert.NotEmpty(t, project)
	assert.Equal(t, "Project1", project.Name)
	assert.Equal(t, "Scenario1", project.Scenarios[0].Name)
	assert.Equal(t, "Case1", project.Scenarios[0].Cases[0].Name)	
	assert.Equal(t, "G1", project.Scenarios[0].Cases[0].Steps[0].Name)	
	assert.Equal(t, "G1.1", project.Scenarios[0].Cases[0].Steps[0].SubSteps[0].Name)	
	assert.Equal(t, "G2", project.Scenarios[0].Cases[0].Steps[1].Name)
	return project
}

func Proto2ProjectTest(t *testing.T, pProject *rpcCli.Project) {
	project := utils.ParseProto2Project(pProject)
	assert.NotEmpty(t, project)
	assert.Equal(t, "Project1", project.Name)
	assert.Equal(t, "Scenario1", project.Scenarios[0].Name)
	assert.Equal(t, "Case1", project.Scenarios[0].Cases[0].Name)	
	assert.Equal(t, "G1", project.Scenarios[0].Cases[0].Steps[0].GetCommon().Name)	
	assert.Equal(t, "G1.1", project.Scenarios[0].Cases[0].Steps[0].GetSubSteps()[0].GetCommon().Name)	
	assert.Equal(t, "G2", project.Scenarios[0].Cases[0].Steps[1].GetCommon().Name)
}

func Step2ProtoTest(t *testing.T) {
	step := getTestStep()
	execution := core.Execution{ID:"1"}
	pStep := utils.ParseStep2Proto(&execution, step)
	assert.NotEmpty(t, pStep)

}

func getTestStep() core.TestStep {
	params := make(map[string]interface{})
	params["URL"] = "https://google.com"
	subStep := steps.NewStep("GET_REQUEST_STEP", "1.1GS", params, nil)
	var subSteps []core.TestStep
	subSteps = append(subSteps, subStep)
	step := steps.NewStep("GET_REQUEST_STEP", "1GS", params, subSteps)
	return step
}

func getTestProjectJson() string {
	json := `{
				"Name": "Project1",
				"Scenarios": [{
					"Name": "Scenario1",
					"Cases": [{
						"Name": "Case1",
						"Steps": [{
								"Name": "G1",
								"Type": "GET_REQUEST_STEP",
								"Parameters": {"URL": "https://google.com"},
								"SubSteps" : [
									{
										"Name": "G1.1",
										"Type": "GET_REQUEST_STEP",
										"Parameters": {"URL": "https://google.com"}
									}
								]
							},
							{
								"Name": "G2",
								"Type": "GET_REQUEST_STEP",
								"Parameters": {"URL": "https://google.com"}
							}
						]
					}]
				}]
			}`

	return json
}
