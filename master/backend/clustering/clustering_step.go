package clustering

import (
	"bitbucket.org/instinctools/gluten/core"
)

var (
	clusteredStepAlias = "CLUSTERED_STEP"
)

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

func (step *ClusteredStep) Run() []core.Metric {
	for node := range GetNodes() {

	}
	return []core.Metric{}
}
