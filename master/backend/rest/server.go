package rest

import (
	"net/http"
	"strconv"

	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func LaunchWebServer(port int) {

	router := httprouter.New()

	//TODO - remove CORS filter ...
	handler := cors.Default().Handler(router)

	router.POST(EXECUTIONS_URL, GetExecution)
	router.GET(RESULTS_URL, GetResults)
	router.POST(STOP_EXECUTION_URL, StopExecution)
	router.GET(NODES_URL, GetNodes)
	router.POST(BUILD_PROJECT_URL, RunProject)
	router.GET(PROJECTS_URL, GetProjects)
	router.GET(EDIT_PROJECT_URL, EditProjectByKey)

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// no implementation
		w.WriteHeader(405)
	})

	logging.WithFields(logging.Fields{
		"error": http.ListenAndServe(":"+strconv.Itoa(port), handler),
	}).Error("Error during starting rpc server")

}
