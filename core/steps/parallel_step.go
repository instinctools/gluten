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
	//validate and preset parameters
	rawThreads := params["THREADS"]
	var resolvedThreads int
	switch rawThreads.(type) {
	case int:
		resolvedThreads = rawThreads.(int)
	default:
		panic("Unsupported parameter type")
	}

	return &ParallelStep{
		CompositeStep{core.BaseTestStep{core.Common{name}, params, substeps}},
		resolvedThreads,
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

func (step *ParallelStep) Run() []core.Metric {
	successRepeats := 0
	for i := 0; i < step.threads; i++ {
		go step.CompositeStep.Run()
		successRepeats++
	}
	return []core.Metric{{Key: "SUCCESS_REPEATS", Val: successRepeats}}
}
