//TODO this file requere refactoring - split into few small parts (one step - one file)
package steps

import (
	"log"
	"net/http"
	"bitbucket.org/instinctools/gluten/core"
)

// TypeComposite & GetRequestStepConstant for global export
const (
	TypeComposite          string = "COMPOSITE"
	typeRequest            string = "REQUEST"
	GetRequestStepConstant        = "GetRequestStep"
	compositeStep                 = "compositeStep"
	repeatStep                    = "repeatStep"
	parallelStep                  = "parallelStep"
)

// Step registry & Step factory
type newStepF func(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep

var stepFactories = make(map[string]newStepF)

func RegisterStepFactory(name string, factory newStepF) {
	if factory == nil {
		panic("Factory does not exist.")
	}
	stepFactories[name] = factory
}

func NewStep(step string, name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
	return stepFactories[step](name, params, substeps)
}

func init() {
	RegisterStepFactory(GetRequestStepConstant, newGetRequestStep)
	RegisterStepFactory(compositeStep, newCompositeStep)
	RegisterStepFactory(repeatStep, newRepeaterStep)
	RegisterStepFactory(parallelStep, newParallelStep)
}

//GetRequestStep ...
type GetRequestStep struct {
	core.BaseTestStep
	url string
}

func newGetRequestStep(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
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
		core.BaseTestStep{core.Common{name}, params, substeps},
		resolvedurl,
	}
}

func (step *GetRequestStep) GetCommon() core.Common {
	return step.Common
}

func (step *GetRequestStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *GetRequestStep) GetSubSteps() []core.TestStep {
	return step.Substeps
}

func (step *GetRequestStep) GetStepType() string {
	return typeRequest
}

func (step *GetRequestStep) BeforeStep() {
	//Do nothing
}

func (step *GetRequestStep) Run() []core.Metric {
	resp, err := http.Get(step.url)
	if err != nil {
		return []core.Metric{{Key: "STATUS", Val: err.Error()}}
	}
	err = resp.Body.Close()
	if err != nil {
		log.Fatal("Error closing", err)
	}
	return []core.Metric{{Key: "STATUS", Val: resp.Status}}
}

//CompositeStep ...
type CompositeStep struct {
	core.BaseTestStep
}

func newCompositeStep(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
	//validate and preset parameters
	return &CompositeStep{
		core.BaseTestStep{core.Common{name}, params, substeps},
	}
}

func (step *CompositeStep) GetCommon() core.Common {
	return step.Common
}

func (step *CompositeStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *CompositeStep) GetSubSteps() []core.TestStep {
	return step.Substeps
}

func (step *CompositeStep) GetStepType() string {
	return TypeComposite
}

func (step *CompositeStep) BeforeStep() {
	//validate and preset parameters
}

func (step *CompositeStep) Run() []core.Metric {
	for _, s := range step.Substeps {
		s.Run()
	}
	return []core.Metric{}
}

//RepeaterStep ...
type RepeaterStep struct {
	CompositeStep
	repeats int
}

func newRepeaterStep(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
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
		CompositeStep{core.BaseTestStep{core.Common{name}, params, substeps}},
		resolvedRepeats,
	}
}

func (step *RepeaterStep) GetCommon() core.Common {
	return step.Common
}

func (step *RepeaterStep) GetParams() map[string]interface{} {
	return step.Parameters
}

func (step *RepeaterStep) GetSubSteps() []core.TestStep {
	return step.Substeps
}

func (step *RepeaterStep) GetStepType() string {
	return TypeComposite
}

func (step *RepeaterStep) BeforeStep() {
}

func (step *RepeaterStep) Run() []core.Metric {
	successRepeats := 0
	for i := 0; i < step.repeats; i++ {
		step.CompositeStep.Run()
		successRepeats++
	}
	return []core.Metric{{Key: "SUCCESS_REPEATS", Val: successRepeats}}
}

//ParallelStep ...
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
	return TypeComposite
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
