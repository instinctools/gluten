package rest

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func StartWebServer(address string) {

	router := httprouter.New()

	//CORS filter ...
	handler := cors.Default().Handler(router)

	router.GET(EXECUTIONS_URL, GetExecution)
	router.GET(RESULTS_URL, GetResults)
	router.POST(STOP_EXECUTION_URL, StopExecution)
	router.POST(EXECUTIONS_URL, StartExecution)

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// not implementation
		w.WriteHeader(405)
	})

	log.Fatal(http.ListenAndServe(address, handler))
}
