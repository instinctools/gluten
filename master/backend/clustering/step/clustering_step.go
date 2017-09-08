package step

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/steps"
	"bitbucket.org/instinctools/gluten/master/backend/ctx"
)

var (
	clusteredStepAlias = "CLUSTERED_STEP"
	nodeStore          = ctx.GlobalContext.NodeStore
)

func init() {
	steps.RegisterStepFactory(clusteredStepAlias, newClusteredStep)
}

type ClusteredStep struct {
	core.BaseTestStep
}

func newClusteredStep(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
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

func (step *ClusteredStep) Run(context *core.Execution, handler core.ResultHandler) {
	nodeStore.SubmitToAll(context, &steps.CompositeStep{
		core.BaseTestStep{
			step.GetCommon(),
			nil,
			step.GetSubSteps(),
		}})
	handler.Handle(core.StepResult{
		ExecutionID: context.ID,
		StepType:    step.GetStepType(),
		Metrics:     []core.Metric{{Key: "NODES", Val: nodeStore.GetNodes()}},
	})
}
