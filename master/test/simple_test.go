package simple_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"bitbucket.org/instinctools/gluten/core/steps"

	assert "github.com/stretchr/testify/require"
)

func TestScenarioWithGetRequests(t *testing.T) {
	case1 := steps.TestCase{
		Name: "Case1",
		Steps: []steps.Step{
			steps.NewGetRequestStep("GStep1", "http://google.com"),
			steps.NewGetRequestStep("GStep2", "http://google.com"),
		}}

	testScenario(case1, t)
}

func TestScenarioWithCompositeStep(t *testing.T) {
	testCase := steps.TestCase{
		Name: "Case1",
		Steps: []steps.Step{
			steps.NewGetRequestStep("GStep1", "http://google.com"),
			steps.NewGetRequestStep("GStep2", "http://google.com"),
			steps.NewCompositeStep("CStep1", []steps.Step{
				steps.NewGetRequestStep("GStep3", "http://google.com"),
				steps.NewGetRequestStep("GStep4", "http://google.com"),
			}),
		}}

	testScenario(testCase, t)
}

func testScenario(testCase steps.TestCase, t *testing.T) {
	originalScenario := steps.Scenario{Name: "Sc1"}
	originalScenario.Add(testCase)
	fmt.Printf("Original scenario: %s\n", testCase)
	originalJson, err := json.Marshal(originalScenario)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Original scenario json: %s\n", string(originalJson))
	deserializedScenario := steps.Scenario{}
	err = json.Unmarshal(originalJson, &deserializedScenario)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deserialized scenario: %s\n", deserializedScenario)
	deserializedScenarioJson, err := json.Marshal(deserializedScenario)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deserialized scenario json: %s\n", string(deserializedScenarioJson))
	assert.Equal(t, originalJson, deserializedScenarioJson)
}
