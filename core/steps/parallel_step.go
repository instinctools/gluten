package steps

import (
	"bitbucket.org/instinctools/gluten/core"
)

var (
	parallelStepAlias = "PARALLEL_STEP"
)

func init() {
	RegisterStepFactory(parallelStepAlias, newParallelStep)
}

type ParallelStep struct {
	CompositeStep
	threads int
}

func newParallelStep(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
	return &ParallelStep{
		CompositeStep{core.BaseTestStep{core.Common{name}, params, substeps}},
		params["THREADS"].(int),
	}
}

func (step *ParallelStep) GetCommon() core.Common {
	return step.Common
}

func (step *ParallelStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *ParallelStep) GetSubSteps() []core.TestStep {
	return step.Substeps
}

func (step *ParallelStep) GetStepType() string {
	return parallelStepAlias
}

func (step *ParallelStep) BeforeStep() {
}

func (step *ParallelStep) Run(context *core.Execution, handler core.ResultHandler) {
	successRepeats := 0
	for i := 0; i < step.threads; i++ {
		go step.CompositeStep.Run(context, handler)
		successRepeats++
	}
	handler.Handle(core.StepResult{
		ExecutionID: context.ID,
		StepType:    step.GetStepType(),
		Metrics:     []core.Metric{{Key: "SUCCESS_REPEATS", Val: successRepeats}},
	})
}
