package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"bitbucket.org/instinctools/gluten/master/backend/ctx"
	"bitbucket.org/instinctools/gluten/shared/utils"
	"encoding/json"
	"strconv"
)

func GetExecution(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(writer, "Error reading request body",
			http.StatusInternalServerError)
	}
	offset, _ := strconv.Atoi(string(body))
	println(offset)
	json.NewEncoder(writer).Encode(ctx.GlobalContext.RawExecutionRepo.Get(8, offset))
	writer.WriteHeader(http.StatusOK)
}

func GetResults(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	id := p.ByName("id")
	json.NewEncoder(writer).Encode(ctx.GlobalContext.RawResultRepo.GetByExecutionId(id, 8, 0))
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

func GetNodes(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(ctx.GlobalContext.NodeStore.GetNodes())
	writer.WriteHeader(http.StatusOK)
}

func RunProject(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(writer, "Error reading request body",
			http.StatusInternalServerError)
	}
	ctx.GlobalContext.ExecutionService.ExecuteProject(utils.ParseProto2Project(utils.DeserializeJsonToProto(string(body))))
	writer.WriteHeader(http.StatusOK)
}

func GetProjects(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(ctx.GlobalContext.ProjectStore.GetProjectNames())
	writer.WriteHeader(http.StatusOK)
}

func EditProjectByKey(writer http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	key := p.ByName("key")
	//TODO: no implementation; need deserialization core.Project to JSON
	println(key)
	writer.WriteHeader(http.StatusOK)
}
