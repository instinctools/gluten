package core_test

import (
	"testing"

	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/steps"
	"github.com/stretchr/testify/mock"
	assert "github.com/stretchr/testify/require"
)

func TestProjectWorkflow(t *testing.T) {
	params := make(map[string]interface{})
	params["URL"] = "http://google.com"
	case1 := core.TestCase{
		Common: core.Common{"Case1"},
		Steps: []core.TestStep{
			steps.NewStep("GET_REQUEST_STEP", "G1", params, nil),
		}}

	scenario1 := core.TestScenario{Common: core.Common{"Sc1"}}
	scenario1.Add(case1)

	project1 := &core.Project{Common: core.Common{"Project1"}}
	project1.Add(scenario1)

	resultHandler := &ResultHandlerMock{}
	resultHandler.On("Handle", mock.Anything).Return().Run(func(args mock.Arguments) {
		stepResult := args.Get(0).(core.StepResult)
		resultHandler.StepResults = append(resultHandler.StepResults, stepResult)
	})
	runner := core.NewDefaultRunner(resultHandler)
	assert.Equal(t, 1, len(project1.GetAllSteps()))
	execution := core.Execution{ID:"1"}
	runner.Run(&execution, project1)

	resultHandler.AssertNumberOfCalls(t, "Handle", 1)
	assert.Equal(t, 1, len(resultHandler.StepResults))

	stepResult := resultHandler.StepResults[0]
	assert.NotEmpty(t, stepResult.ExecutionID)
	assert.Equal(t, "Completed", stepResult.Status)
	assert.Equal(t, case1.Steps[0].GetStepType(), stepResult.StepType)

	assert.Equal(t, 1, len(stepResult.Metrics))
	metric := stepResult.Metrics[0]
	assert.Equal(t, "STATUS", metric.Key)

}

type ResultHandlerMock struct {
	mock.Mock
	StepResults []core.StepResult
}

func (m *ResultHandlerMock) Handle(result core.StepResult) {
	m.Called(result)
}
