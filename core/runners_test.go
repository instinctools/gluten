package core_test

import (
	"testing"

	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/steps"
	"github.com/stretchr/testify/mock"
	assert "github.com/stretchr/testify/require"
)

func TestProjectWorkflow(t *testing.T) {
	case1 := core.TestCase{
		Name: "Case1",
		Steps: []core.Step{
			steps.NewGetRequestStep("Step1", "http://google.com"),
		}}

	scenario1 := core.Scenario{Name: "Sc1"}
	scenario1.Add(case1)

	project1 := &core.Project{Name: "Project1"}
	project1.Add(scenario1)

	resultHandler := &ResultHandlerMock{}
	resultHandler.On("Handle", mock.Anything).Return().Run(func(args mock.Arguments) {
		stepResult := args.Get(0).(core.StepResult)
		resultHandler.StepResults = append(resultHandler.StepResults, stepResult)
	})
	runner := core.DefaultRunner{
		Handler: resultHandler,
	}
	assert.Equal(t, 1, len(project1.GetSteps()))
	runner.Run(project1)

	resultHandler.AssertNumberOfCalls(t, "Handle", 1)
	assert.Equal(t, 1, len(resultHandler.StepResults))

	stepResult := resultHandler.StepResults[0]
	assert.NotEmpty(t, stepResult.ExecutionID)
	assert.Equal(t, "OK", stepResult.Status)
	assert.NoError(t, stepResult.Error)
	assert.Equal(t, case1.Steps[0], stepResult.Step)

	assert.Equal(t, 1, len(stepResult.Metrics))
	metric := stepResult.Metrics[0]
	assert.Equal(t, "TIME", metric.Key)
	assert.Equal(t, 100, metric.Val)

}

type ResultHandlerMock struct {
	mock.Mock
	StepResults []core.StepResult
}

func (m *ResultHandlerMock) Handle(result core.StepResult) {
	m.Called(result)
}
