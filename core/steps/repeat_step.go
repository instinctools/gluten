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
	//validate and preset parameters
	rawRepeats := params["REPEATS"]
	var resolvedRepeats int
	switch rawRepeats.(type) {
	case int:
		resolvedRepeats = rawRepeats.(int)
	default:
		panic("Unsupported parameter type")
	}

	return &RepeatStep{
		CompositeStep{core.BaseTestStep{core.Common{name}, params, substeps}},
		resolvedRepeats,
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

func (step *RepeatStep) Run() []core.Metric {
	successRepeats := 0
	for i := 0; i < step.repeats; i++ {
		step.CompositeStep.Run()
		successRepeats++
	}
	return []core.Metric{{Key: "SUCCESS_REPEATS", Val: successRepeats}}
}
