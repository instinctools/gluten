package steps

import (
	"encoding/json"
)

//CompositeStep ...
type CompositeStep struct {
	BaseStep
	SubSteps []Step
}

func NewCompositeStep(name string, subSteps []Step) *CompositeStep {
	return &CompositeStep{
		BaseStep{name},
		subSteps,
	}
}

func (step *CompositeStep) Run() []StepResult {
	stepResults := []StepResult{}
	for _, s := range step.GetSubSteps() {
		stepResults = append(stepResults, s.Run()...)
	}
	return stepResults
}

func (step *CompositeStep) BeforeStep() {
	//validate and preset parameters
}

func (step *CompositeStep) GetSubSteps() []Step {
	return step.SubSteps
}

func (step *CompositeStep) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"type":     "Composite",
		"name":     step.GetName(),
		"subSteps": step.GetSubSteps(),
	})
}

func (step *CompositeStep) UnmarshalJSON(data []byte) error {
	var tmp map[string]interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	step.Name = tmp["name"].(string)

	//TODO make me more clear as in test_case_deserializer.go
	if subSteps, ok := tmp["subSteps"]; ok {
		for _, subStep := range subSteps.([]interface{}) {
			stepType := subStep.(map[string]interface{})["type"]
			if stepType == "GetRequest" {
				var getRequest = &GetRequestStep{}
				//TODO handle error
				tmpJson, _ := json.Marshal(subStep)
				//TODO handle error
				json.Unmarshal(tmpJson, getRequest)
				step.SubSteps = append(step.SubSteps, getRequest)
			}
		}
	}
	return nil
}
