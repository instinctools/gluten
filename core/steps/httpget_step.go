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
	return httpgetStepAlias
}

func (step *GetRequestStep) BeforeStep() {
	//Do nothing
}

func (step *GetRequestStep) Run(context *core.Execution) []core.Metric {
	resp, err := http.Get(step.url)
	if err != nil {
		return []core.Metric{{Key: "STATUS", Val: err.Error()}}
	}
	err = resp.Body.Close()
	if err != nil {
		logging.WithFields(logging.Fields{
			"error": err,
		}).Error("Error during closing connection")
	}
	return []core.Metric{{Key: "STATUS", Val: resp.Status}}
}
