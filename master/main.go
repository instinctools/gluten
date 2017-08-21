package main

import (
	"log"
	"net/http"

	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

const BASE_URL string = "/api/"

type Message struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
}

var Result struct {
	id   uint
	text string
}

func main() {

	router := httprouter.New()

	//CORS filter ...
	handler := cors.Default().Handler(router)

	router.GET(BASE_URL+"executions/", GetExecution)
	router.GET(BASE_URL+"executions/:id/results/", GetResults)
	router.POST(BASE_URL+"executions/:id/stop/", StopExecution)
	router.POST(BASE_URL+"executions/", StartExecution)

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// not implementation
		w.WriteHeader(405)
	})

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func generateResults() []ExecutionResult {
	return []ExecutionResult{
		{
			ID:      12321323,
			Created: time.Now().UnixNano(),
			Metrics: []Metric{
				{ExecutionResultID: 335669, Key: "PUT", Value: "https://facebook.com"},
			},
		},
		{
			ID:      989565,
			Created: time.Now().UnixNano(),
			Metrics: []Metric{
				{ExecutionResultID: 4111, Key: "GET", Value: "http://google.com"},
			},
		},
		{
			ID:      113,
			Created: time.Now().UnixNano(),
			Metrics: []Metric{
				{ExecutionResultID: 986, Key: "GET", Value: "http://google.com"},
			},
		},
	}
}

func generateExecutions() []Execution {
	return []Execution{
		{
			ID:         12312,
			Created:    time.Now().UnixNano(),
			Parameters: "smth",
			Result: ExecutionResult{
				ID:      12321323,
				Created: time.Now().UnixNano(),
				Metrics: []Metric{
					{ExecutionResultID: 335669, Key: "PUT", Value: "https://facebook.com"},
				},
			},
			ResultID: 007,
		},
		{
			ID:         56436546,
			Created:    time.Now().UnixNano(),
			Parameters: "hello dude",
			Result: ExecutionResult{
				ID:      989565,
				Created: time.Now().UnixNano(),
				Metrics: []Metric{
					{ExecutionResultID: 4111, Key: "GET", Value: "http://google.com"},
				},
			},
			ResultID: 745,
		},
		{
			ID:         19893,
			Created:    time.Now().UnixNano(),
			Parameters: "bye dude",
			Result: ExecutionResult{
				ID:      113,
				Created: time.Now().UnixNano(),
				Metrics: []Metric{
					{ExecutionResultID: 986, Key: "GET", Value: "http://google.com"},
				},
			},
			ResultID: 888,
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
