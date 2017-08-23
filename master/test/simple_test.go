package test

import (
	"bitbucket.org/instinctools/gluten/core/steps"
	"encoding/json"
	"fmt"
	"testing"
)

type Model struct {
	Case steps.TestCase
}

func TestRunner(t *testing.T) {
	case1 := steps.TestCase{
		Name: "Case1",
		Steps: []steps.Step{
			steps.NewGetRequestStep("GStep1", "http://google.com"),
			steps.NewGetRequestStep("GStep2", "http://google.com"),
		}}

	fmt.Print("Case obj - ")
	fmt.Println(case1)

	scenario1 := steps.Scenario{Name: "Sc1"}
	scenario1.Add(case1)

	jsonS, err := json.Marshal(scenario1)
	if err != nil {
		fmt.Println(err)
	}
	str := string(jsonS)
	fmt.Print("Scenario json - ")
	fmt.Println(str)
	scenario2 := steps.Scenario{}
	err = json.Unmarshal([]byte(jsonS), &scenario2)
	if err != nil {
		fmt.Print("48 - ")
		fmt.Println(err)
	}
	fmt.Print("Case after unmarsh - ")
	fmt.Println(scenario2)

	jsonS, err = json.Marshal(scenario2)
	if err != nil {
		fmt.Println(err)
	}
	str = string(jsonS)
	fmt.Print("Case unmarshed json - ")
	fmt.Println(str)

}
