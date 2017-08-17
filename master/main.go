package main

import (
	"net/http"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

const BASE_URL string = "/api/"

func GetExecution(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(generateExecutions())
	writer.WriteHeader(200)
}

func Hello(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(struct {
		ID     string
		Status string
	}{
		ID:     "1",
		Status: "GET",
	})
	writer.WriteHeader(200)
}

func main() {
	router := httprouter.New()

	//CORS filter ...
	handler := cors.Default().Handler(router)

	router.GET(BASE_URL+"executions/", GetExecution)
	router.GET(BASE_URL, Hello)
	router.POST(BASE_URL+"stop/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Header().Set("Content-Type", "application/json")
		println(request.Body)
		writer.WriteHeader(201)
	})

	http.ListenAndServe(":8080", handler)
}

func generateExecutions() []Execution {
	return []Execution{
		{
			ID:         12312,
			Created:    3,
			Parameters: "smth",
			Result: ExecutionResult{
				ID:      12321323,
				Created: 4,
				Metrics: []Metric{
					{ExecutionResultID: 335669, Key: "PUT", Value: "https://facebook.com"},
				},
			},
			ResultID: 007,
		},
		{
			ID:         56436546,
			Created:    1,
			Parameters: "hello dude",
			Result: ExecutionResult{
				ID:      989565,
				Created: 2,
				Metrics: []Metric{
					{ExecutionResultID: 4111, Key: "GET", Value: "http://google.com"},
				},
			},
			ResultID: 745,
		},
	}
}

type Execution struct {
	ID         uint `gorm:"primary_key"`
	Created    int64
	Parameters string
	Result     ExecutionResult `gorm:"ForeignKey:ResultID"`
	ResultID   uint
}

type ExecutionResult struct {
	ID      uint `gorm:"primary_key"`
	Created int64
	Metrics []Metric
}

type Metric struct {
	ExecutionResultID uint
	Key               string
	Value             string
}
