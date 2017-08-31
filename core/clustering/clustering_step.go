package clustering

import (
	"bitbucket.org/instinctools/gluten/core"
)

var (
	clusteredStepAlias = "CLUSTERED_STEP"
)

type ClusteredStep struct {
	core.BaseTestStep
	clusterContext *ClusterContext
}

func newRepeatStep(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
	//validate and preset parameters
	rawContext := params["CLUSTERCONTEXT"]
	var resolvedContext ClusterContext
	switch rawContext.(type) {
	case ClusterContext:
		resolvedContext = rawContext.(ClusterContext)
	default:
		panic("Unsupported parameter type")
	}

	return &ClusteredStep{
		core.BaseTestStep{core.Common{name}, params, substeps},
		&resolvedContext,
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
	//	for _, s := range step.clusterContext.GetNodes() {
	//		s.SubmitAndExecute()
	//	}
	return []core.Metric{}
}
