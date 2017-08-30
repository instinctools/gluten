package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadJSONFile(pathToFile string) string {
	plan, _ := ioutil.ReadFile(pathToFile)
	data := Project{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		log.Println("File not found", err)
	}
	str, err := json.Marshal(data)
	if err != nil {
		log.Println("Bad convert", err)
	}
	return string(str)
}

func AutoGenerateConfig(filename string) {
	//TODO - generateJSON should be a string, don't need to marshal it to string
	generateJSON := Project{
	//			Name: "1",
	//			Scenarios: []TestScenario{
	//				{Name: "2", Cases: []TestCase{
	//					{Name: "3", Steps: []TestStep{
	//						{
	//							Type: "GETRequest",
	//							Params: []Params{
	//								{
	//									Key:   "URL",
	//									Value: "https://google.com",
	//								},
	//							},
	//						},
	//					}},
	//				}},
	//			},
	}
	response, err := json.Marshal(generateJSON)
	if err != nil {
		log.Fatal("Bad convert", err)
	}
	err = ioutil.WriteFile(filename, response, 0644)
	if err != nil {
		log.Fatal("Fatal error for writing", err)
	}
}
