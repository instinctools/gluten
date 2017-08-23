package backend

import (
	"log"
	"net/http"

	controller "bitbucket.org/instinctools/gluten/master/backend/controllers"
	route "bitbucket.org/instinctools/gluten/master/backend/routes"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func StartWebServer(address string) {

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

	log.Fatal(http.ListenAndServe(address, handler))
}
