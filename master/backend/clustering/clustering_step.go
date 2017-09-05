package clustering

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/steps"
)

var (
	clusteredStepAlias = "CLUSTERED_STEP"
)

func init() {
	steps.RegisterStepFactory(clusteredStepAlias, newClusteredStep)
}

type ClusteredStep struct {
	core.BaseTestStep
}

func newClusteredStep(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
	//validate and preset parameters

	return &ClusteredStep{
		core.BaseTestStep{core.Common{name}, params, substeps},
	}
}

func (step *ClusteredStep) GetCommon() core.Common {
	return step.Common
}

func (step *ClusteredStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *ClusteredStep) GetSubSteps() []core.TestStep {
	return step.Substeps
}

func (step *ClusteredStep) GetStepType() string {
	return clusteredStepAlias
}

func (step *ClusteredStep) BeforeStep() {
	//validate and preset parameters
}

func (step *ClusteredStep) Run(context *core.Execution) []core.Metric {
	for _, node := range GetNodes() {
		SubmitOverRPC(node, context, &steps.CompositeStep{step.BaseTestStep})
	}
	return nil
}
