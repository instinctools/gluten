package clustering

import (
	core "bitbucket.org/instinctools/gluten/core"
)

const CLUSTER_WRAPPER_STEP = "CLUSTER_WRAPPER_STEP"

func init() {
    core.RegisterStepFactory(CLUSTER_WRAPPER_STEP, newClusteringWrapperStep)
}

//ClusteringWrapperStep
type ClusteringWrapperStep struct {
	core.CompositeStep
	cluster *ClusterContext
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
	
	return &ClusteringWrapperStep{
		core.CompositeStep{core.BaseTestStep{core.Common{name}, params, substeps}},
		&resolvedcluster,
		substeps[0],
	}
}

func (step *ClusteringWrapperStep) GetCommon() core.Common {
	return step.Common
}

func (step *ClusteringWrapperStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *ClusteringWrapperStep) GetSubSteps() []core.TestStep {
	return step.Substeps
}

func (step *ClusteringWrapperStep) GetStepType() string {
	return core.TYPE_COMPOSITE
}

func (step *ClusteringWrapperStep) BeforeStep() {
	//validate and preset parameters
}

func (step *ClusteringWrapperStep) Run() []core.Metric {
	for _, s := range step.cluster.GetNodes() {
		s.SubmitAndExecute(step.wrappedStep)
	}
	return []core.Metric{}
}

//Clustering models
type ClusterNode struct {
	url string
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

