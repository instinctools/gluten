package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"encoding/json"

	node "bitbucket.org/instinctools/gluten/master/backend/clustering"
	"bitbucket.org/instinctools/gluten/shared/persistence/gorm"
)

func GetExecution(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(gorm.RawExecutionsRepoInstance.Get(10, 0))
	writer.WriteHeader(http.StatusOK)
}

func GetResults(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id := p.ByName("id")
	json.NewEncoder(writer).Encode(gorm.RawResultsRepoInstance.GetByExecutionId(id, 10, 0))
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

func GetNodes(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(node.GetNodes())
	writer.WriteHeader(http.StatusOK)
}
