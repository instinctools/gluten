package rest

import (
	"bitbucket.org/instinctools/gluten/shared/logging"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
	"strconv"
)

func LaunchWebServer(port int) {

	router := httprouter.New()

	//TODO - remove CORS filter ...
	handler := cors.Default().Handler(router)

	router.GET(EXECUTIONS_URL, GetExecution)
	router.GET(RESULTS_URL, GetResults)
	router.POST(STOP_EXECUTION_URL, StopExecution)
	router.POST(EXECUTIONS_URL, StartExecution)

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// not implementation
		w.WriteHeader(405)
	})

	logging.WithFields(logging.Fields{
		"error": http.ListenAndServe(":"+strconv.Itoa(port), handler),
	}).Error("Error during starting rpc server")

}
