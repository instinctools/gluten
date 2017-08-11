package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//for example
const default_path string = "/home/INSTINCTOOLS/[username]/go-tutorial/src/cli/file.json"

func ReadJsonFile(pathToFile string) string {
	plan, _ := ioutil.ReadFile(pathToFile)
	data := JsonModel{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		log.Printf("File not found", err)
	}
	str, err := json.Marshal(data)
	if err != nil {
		log.Printf("Bad convert", err)
	}
	return string(str)
}

func AutoGenerateConfig(filename string) {
	generateJson := JsonModel{
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
	response, err := json.Marshal(generateJson)
	if err != nil {
		log.Printf("Bad convert", err)
	}
	ioutil.WriteFile(filename, response, 0644)
}
