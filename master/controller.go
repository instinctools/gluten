package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GetExecution(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(generateExecutions())
	writer.WriteHeader(200)
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
