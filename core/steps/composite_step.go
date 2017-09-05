package steps

import (
	"bitbucket.org/instinctools/gluten/core"
)

var (
	compositeStepAlias = "COMPOSITE_STEP"
)

func init() {
	RegisterStepFactory(compositeStepAlias, newCompositeStep)
}

type CompositeStep struct {
	core.BaseTestStep
}

func newCompositeStep(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
	//validate and preset parameters
	return &CompositeStep{
		core.BaseTestStep{core.Common{name}, params, substeps},
	}
}

func (step *CompositeStep) GetCommon() core.Common {
	return step.Common
}

func (step *CompositeStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *CompositeStep) GetSubSteps() []core.TestStep {
	return step.Substeps
}

func (step *CompositeStep) GetStepType() string {
	return compositeStepAlias
}

func (step *CompositeStep) BeforeStep() {
	//validate and preset parameters
}

func (step *CompositeStep) Run(context *core.Execution) []core.Metric {
	for _, s := range step.Substeps {
		s.Run(context)
	}
	return []core.Metric{}
}
