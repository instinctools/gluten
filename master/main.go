package main

import (
	"log"
	"net/http"

	controller "bitbucket.org/instinctools/gluten/master/controllers"
	route "bitbucket.org/instinctools/gluten/master/routes"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {

	router := httprouter.New()

	//CORS filter ...
	handler := cors.Default().Handler(router)

	router.GET(route.EXECUTIONS_URL, controller.GetExecution)
	router.GET(route.RESULTS_URL, controller.GetResults)
	router.POST(route.STOP_EXECUTION_URL, controller.StopExecution)
	router.POST(route.EXECUTIONS_URL, controller.StartExecution)

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// not implementation
		w.WriteHeader(405)
	})

	log.Fatal(http.ListenAndServe(":8080", handler))
}
