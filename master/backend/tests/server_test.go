package tests

import (
	"net/http"

	"testing"

	"net/http/httptest"

	controller "bitbucket.org/instinctools/gluten/master/backend/controllers"
	route "bitbucket.org/instinctools/gluten/master/backend/routes"

	"github.com/julienschmidt/httprouter"
	assert "github.com/stretchr/testify/require"
)

func TestRequests(t *testing.T) {

	recorder := doGet(route.EXECUTIONS_URL, controller.GetExecution)
	assert.Equal(t, http.StatusOK, recorder.Code)

	recorder = doGet(route.RESULTS_URL, controller.GetResults)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func doGet(path string, handle func(writer http.ResponseWriter, r *http.Request, p httprouter.Params)) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", path, nil)
	router := httprouter.New()
	router.Handle("GET", path, handle)
	router.ServeHTTP(recorder, request)
	return recorder
}
