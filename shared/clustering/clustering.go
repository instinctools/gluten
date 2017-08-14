package clustering

import (
	"bitbucket.org/instinctools/gluten/core"
)

const (
	clusterWrapperStep = "clusterWrapperStep"
)

func init() {
	core.RegisterStepFactory(clusterWrapperStep, newClusteringWrapperStep)
}

//ClusterWrapperStep ...
type ClusterWrapperStep struct {
	core.CompositeStep
	cluster     *ClusterContext
	wrappedStep core.TestStep
}

func newClusteringWrapperStep(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
	//validate and preset parameters
	rawCluster := params["URL"]
	var resolvedcluster ClusterContext
	switch rawCluster.(type) {
	case string:
		resolvedcluster = rawCluster.(ClusterContext)
	default:
		panic("Unsupported parameter type")
	}
	if len(substeps) != 1 {
		panic("Number of sub steps can't be different then 1")
	}

	return &ClusterWrapperStep{
		core.CompositeStep{core.BaseTestStep{core.Common{name}, params, substeps}},
		&resolvedcluster,
		substeps[0],
	}
}

func (step *ClusterWrapperStep) GetCommon() core.Common {
	return step.Common
}

func (step *ClusterWrapperStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *ClusterWrapperStep) GetSubSteps() []core.TestStep {
	return step.Substeps
}

func (step *ClusterWrapperStep) GetStepType() string {
	return core.TypeComposite
}

func (step *ClusterWrapperStep) BeforeStep() {
	//validate and preset parameters
}

func (step *ClusterWrapperStep) Run() []core.Metric {
	for _, s := range step.cluster.GetNodes() {
		s.SubmitAndExecute(step.wrappedStep)
	}
	return []core.Metric{}
}

//ClusterNode ...
type ClusterNode struct {
	//url string
}

func (node *ClusterNode) SubmitAndExecute(step core.TestStep) {
	//TODO - RPC impl here
}

type ClusterContext struct {
	nodes []ClusterNode
}

func (context *ClusterContext) GetNodes() []ClusterNode {
	return context.nodes

}
