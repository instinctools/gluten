package steps

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/shared/logging"
	"net/http"
)

var (
	httpgetStepAlias = "GET_REQUEST_STEP"
)

func init() {
	RegisterStepFactory(httpgetStepAlias, newGetRequestStep)
}

type GetRequestStep struct {
	core.BaseTestStep
	url string
}

func newGetRequestStep(name string, params map[string]interface{}, substeps []core.TestStep) core.TestStep {
	return &GetRequestStep{
		core.BaseTestStep{core.Common{name}, params, substeps},
		params["URL"].(string),
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
	return httpgetStepAlias
}

func (step *GetRequestStep) BeforeStep() {
	//Do nothing
}

func (step *GetRequestStep) Run(context *core.Execution, handler core.ResultHandler) {
	resp, err := http.Get(step.url)
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error during sending request")
	}
	err = resp.Body.Close()
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error during closing connection")
	}
	handler.Handle(core.StepResult{
		ExecutionID: context.ID,
		Status:      "COMPLETED",
		StepType:    step.GetStepType(),
		Metrics:     []core.Metric{{Key: "RESPONSE_STATUS", Val: resp.Status}},
	})
}
