package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"time"

	"github.com/julienschmidt/httprouter"
)

func GetExecution(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(generateExecutions())
	writer.WriteHeader(http.StatusOK)
}

func GetResults(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(generateResults())
	writer.WriteHeader(200)
}

func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(struct {
		ID     string
		Status string
	}{
		ID:     "1",
		Status: "GET",
	})
	w.WriteHeader(200)
}

func StopExecution(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := p.ByName("id")
	fmt.Fprint(w, "POST done")

	fmt.Println("Stop this execution: " + id)
	w.WriteHeader(201)
}

func StartExecution(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	fmt.Fprint(w, "POST done")

	//submit current execution and start him

	fmt.Println(string(body))
	w.WriteHeader(201)
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
