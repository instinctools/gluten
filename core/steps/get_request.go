package steps

import (
	"net/http"

	"bitbucket.org/instinctools/gluten/core"
)

//GetRequestStep ...
type GetRequestStep struct {
	core.BaseStep
	Url string
}

func NewGetRequestStep(name string, url string) *GetRequestStep {
	return &GetRequestStep{
		core.BaseStep{Name: name},
		url,
	}
}

func (step *GetRequestStep) Run() []core.StepResult {
	//TODO measure time
	time := 100
	resp, err := http.Get(step.Url)
	if err == nil {
		err = resp.Body.Close()

	}
	if err != nil {
		return []core.StepResult{{Status: "NOT_OK", Step: step, Error: err}}
	}
	return []core.StepResult{{
		Status:  "OK",
		Step:    step,
		Metrics: []core.Metric{{Key: "TIME", Val: time}},
	}}
}
