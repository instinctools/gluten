package steps

import (
	"bitbucket.org/instinctools/gluten/core"
	"net/http"
)

//GetRequestStep ...
type GetRequestStep struct {
	core.BaseSingleStep
}

func NewGetRequestStep(name string, url string) *GetRequestStep {
	return &GetRequestStep{
		core.BaseSingleStep{
			core.BaseStep{name},
			url,
		},
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
