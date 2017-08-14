package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadJSONFile(pathToFile string) string {
	plan, _ := ioutil.ReadFile(pathToFile)
	data := JSONModel{}
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
	generateJSON := JSONModel{
		Project{
			Name: "1",
			TestSuits: []TestSuite{
				{Name: "2", TestCases: []TestCase{
					{Name: "3", TestSteps: []TestStep{
						{
							Name: "4",
							Type: "GETRequest",
							Params: []Params{
								{
									Key:   "URL",
									Value: "https://google.com",
								},
							},
						},
					}},
				}},
			},
		},
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
