package steps

import (
	"encoding/json"
	"net/http"
)

//GetRequestStep ...
type GetRequestStep struct {
	BaseStep
	Url string
}

func NewGetRequestStep(name string, url string) *GetRequestStep {
	return &GetRequestStep{
		BaseStep{Name: name},
		url,
	}
}

func (step *GetRequestStep) Run() []StepResult {
	//TODO measure time
	time := 100
	resp, err := http.Get(step.Url)
	if err == nil {
		err = resp.Body.Close()

	}
	if err != nil {
		return []StepResult{{Status: "NOT_OK", Step: step, Error: err}}
	}
	return []StepResult{{
		Status:  "OK",
		Step:    step,
		Metrics: []Metric{{Key: "TIME", Val: time}},
	}}
}

func (step *GetRequestStep) GetName() string {
	return step.Name
}

func (step *GetRequestStep) BeforeStep() {
}

func (g *GetRequestStep) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]string{
		"type": "GetRequest",
		"name": g.GetName(),
		"url":  g.Url,
	})
}
