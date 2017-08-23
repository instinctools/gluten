package test

import (
	"fmt"
	"reflect"
	"testing"
)

type Case struct {
	Steps []Step
}

func (c *Case) GetSteps() []Step {
	return c.Steps
}

type Step interface {
	GetName() string
}

type SingleStep interface {
	GetUrl() string
}

type MultipleStep interface {
	GetSubSteps() []SingleStep
}

type BaseSingleStep struct {
	Name string
	Url  string
}

func (st *BaseSingleStep) GetName() string {
	return st.Name
}

func (st *BaseSingleStep) GetUrl() string {
	return st.Url
}

type GetRequestStep struct {
	BaseSingleStep
	Message string
}

func (st *GetRequestStep) ShowMessage() {
	fmt.Println(st.Message)
}

type BaseMultipleStep struct {
	Name     string
	SubSteps []SingleStep
}

func (st *BaseMultipleStep) GetName() string {
	return st.Name
}

func (st *BaseMultipleStep) GetSubSteps() []SingleStep {
	return st.SubSteps
}

type CompositeStep struct {
	BaseMultipleStep
	Message string
}

func (st *CompositeStep) ShowMessage() {
	fmt.Println(st.Message)
}

func TestRunner(t *testing.T) {
	fmt.Println("TestStarted")

	c := Case{Steps: []Step{
		&GetRequestStep{BaseSingleStep{Name: "Get1", Url: "URL"}, "Message"},
		&GetRequestStep{BaseSingleStep{Name: "Get2", Url: "URL"}, "Message"},
		&CompositeStep{BaseMultipleStep{Name: "Composite1", SubSteps: []SingleStep{
			&GetRequestStep{BaseSingleStep{Name: "Get3", Url: "URL"}, "Message"},
			&GetRequestStep{BaseSingleStep{Name: "Get4", Url: "URL"}, "Message"},
			},}, "Message"},
	}}
	step := c.GetSteps()[0]


	//BaseSingleStep(step).GetName()

	fmt.Println(reflect.TypeOf(step))
}
