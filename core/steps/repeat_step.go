package steps

import (
	"bitbucket.org/instinctools/gluten/core"
)

var (
	repeatStepAlias = "REPEAT_STEP"
)

func init() {
	RegisterStepFactory(repeatStepAlias, newRepeatStep)
}

type RepeatStep struct {
	CompositeStep
	repeats int
}

func newRepeatStep(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
	return &RepeatStep{
		CompositeStep{core.BaseTestStep{core.Common{name}, params, substeps}},
		params["REPEATS"].(int),
	}
}

func (step *RepeatStep) GetCommon() core.Common {
	return step.Common
}

func (step *RepeatStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *RepeatStep) GetSubSteps() []core.TestStep {
	return step.Substeps
}

func (step *RepeatStep) GetStepType() string {
	return repeatStepAlias
}

func (step *RepeatStep) BeforeStep() {
}

func (step *RepeatStep) Run(context *core.Execution, handler core.ResultHandler) {
	successRepeats := 0
	for i := 0; i < step.repeats; i++ {
		step.CompositeStep.Run(context, handler)
		successRepeats++
	}
	handler.Handle(core.StepResult{
		ExecutionID: context.ID,
		StepType:    step.GetStepType(),
		Metrics:     []core.Metric{{Key: "SUCCESS_REPEATS", Val: successRepeats}},
	})
}
