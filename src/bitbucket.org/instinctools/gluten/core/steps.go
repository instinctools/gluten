package core

import (
	"net/http"
)

const TYPE_REQUEST string = "REQUEST"
const TYPE_COMPOSITE string = "COMPOSITE"

//  Requests steps
//	GetRequestStep
type GetRequestStep struct {
	BaseTestStep
	url string
}

func newGetRequestStep(name string, params map[string]interface{}, substeps []TestStep) *GetRequestStep {
	//validate and preset parameters
	url := params["URL"]
	resolvedurl := ""
	switch url.(type) {
	case string:
		resolvedurl = url.(string)
	default:
		panic("Unsupported parameter type")
	}

	return &GetRequestStep{
		BaseTestStep{Common{name}, params, substeps},
		resolvedurl,
	}
}

func (step *GetRequestStep) GetCommon() Common {
	return step.Common
}

func (step *GetRequestStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *GetRequestStep) GetSubSteps() []TestStep {
	return step.Substeps
}

func (step *GetRequestStep) GetStepType() string {
	return TYPE_REQUEST
}

func (step *GetRequestStep) BeforeStep() {
}

func (step *GetRequestStep) Run() []Metric {
	resp, err := http.Get(step.url)
	if err != nil {
		return []Metric{Metric{Key: "STATUS", Val: err.Error()}}
	}
	resp.Body.Close()
	return []Metric{Metric{Key: "STATUS", Val: resp.Status}}
}

// Composite steps
//	CompositeStep
type CompositeStep struct {
	BaseTestStep
}

func newCompositeStep(name string, params map[string]interface{}, substeps []TestStep) *CompositeStep {
	//validate and preset parameters
	return &CompositeStep{
		BaseTestStep{Common{name}, params, substeps},
	}
}

func (step *CompositeStep) GetCommon() Common {
	return step.Common
}

func (step *CompositeStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *CompositeStep) GetSubSteps() []TestStep {
	return step.Substeps
}

func (step *CompositeStep) GetStepType() string {
	return TYPE_COMPOSITE
}

func (step *CompositeStep) BeforeStep() {
	//validate and preset parameters
}

func (step *CompositeStep) Run() []Metric {
	for _, s := range step.Substeps {
		s.Run()
	}
	return []Metric{}
}

//	RepeaterStep
type RepeaterStep struct {
	CompositeStep
	repeats int
}

func newRepeaterStep(name string, params map[string]interface{}, substeps []TestStep) *RepeaterStep {
	//validate and preset parameters
	rawRepeats := params["REPEATS"]
	var resolvedRepeats int
	switch rawRepeats.(type) {
	case int:
		resolvedRepeats = rawRepeats.(int)
	default:
		panic("Unsupported parameter type")
	}

	return &RepeaterStep{
		CompositeStep{BaseTestStep{Common{name}, params, substeps}},
		resolvedRepeats,
	}
}

func (step *RepeaterStep) GetCommon() Common {
	return step.Common
}

func (step *RepeaterStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *RepeaterStep) GetSubSteps() []TestStep {
	return step.Substeps
}

func (step *RepeaterStep) GetStepType() string {
	return TYPE_COMPOSITE
}

func (step *RepeaterStep) BeforeStep() {
}

func (step *RepeaterStep) Run() []Metric {
	success_repeats := 0
	for i := 0; i < step.repeats; i++ {
		step.CompositeStep.Run()
		success_repeats++
	}
	return []Metric{Metric{Key: "SUCCESS_REPEATS", Val: success_repeats}}
}

//	ParallelStep
type ParallelStep struct {
	CompositeStep
	threads int
}

func newParallelStep(name string, params map[string]interface{}, substeps []TestStep) *ParallelStep {
	//validate and preset parameters
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
		CompositeStep{BaseTestStep{Common{name}, params, substeps}},
		resolvedThreads,
	}
}

func (step *ParallelStep) GetCommon() Common {
	return step.Common
}

func (step *ParallelStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *ParallelStep) GetSubSteps() []TestStep {
	return step.Substeps
}

func (step *ParallelStep) GetStepType() string {
	return TYPE_COMPOSITE
}

func (step *ParallelStep) BeforeStep() {
}

func (step *ParallelStep) Run() []Metric {
	success_repeats := 0
	for i := 0; i < step.threads; i++ {
		go func() {
			step.CompositeStep.Run()
		}()
		success_repeats++
	}
	return []Metric{Metric{Key: "SUCCESS_REPEATS", Val: success_repeats}}
}
