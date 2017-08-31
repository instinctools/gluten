package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"strconv"

	log "bitbucket.org/instinctools/gluten/shared/logging"
	repo "bitbucket.org/instinctools/gluten/shared/persistence/repository"
)

func GetExecution(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(repo.GetExecutions())
	writer.WriteHeader(http.StatusOK)
}

func GetResults(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil && uint(id) != 0 {
		log.WithFields(log.Fields{
			"id": id,
		}).Fatal("Error convert")
	}
	json.NewEncoder(writer).Encode(repo.GetResults(uint(id)))
	writer.WriteHeader(200)
}

func StopExecution(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := p.ByName("id")
	fmt.Fprint(w, "POST done")

	// no implementation
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

	// no implementation
	//submit current execution and start him

	fmt.Println(string(body))
	w.WriteHeader(201)
}
