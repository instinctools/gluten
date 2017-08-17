package clustering

import (
	"bitbucket.org/instinctools/gluten/core"
)

//ClusterWrapperStep ...
type ClusteredStep struct {
	core.BaseStep
	clusterContext *ClusterContext
	delegate       core.Step
}

func NewClusteredStep(name string, clusterContext ClusterContext, delegate core.Step) *ClusteredStep {
	return &ClusteredStep{
		core.BaseStep{Name: name},
		&clusterContext,
		delegate,
	}
}

func (step *ClusteredStep) BeforeStep() {
	//validate and preset parameters
}

func (step *ClusteredStep) Run() []core.StepResult {
	for _, s := range step.clusterContext.GetNodes() {
		s.SubmitAndExecute(step.delegate)
	}
	return []core.StepResult{}
}

//ClusterNode ...
type ClusterNode struct {
	//url string
}

func (node *ClusterNode) SubmitAndExecute(step core.Step) {
	//TODO - RPC impl here
}

type ClusterContext struct {
	nodes []ClusterNode
}

func (context *ClusterContext) GetNodes() []ClusterNode {
	return context.nodes

}
