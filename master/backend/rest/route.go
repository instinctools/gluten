package rest

const base_url string = "/api/"

var (
	EXECUTIONS_URL string = base_url + "executions/"

	RESULTS_URL string = EXECUTIONS_URL + ":id/results/"

	STOP_EXECUTION_URL string = EXECUTIONS_URL + ":id/stop/"

	AGGREGATE_STATISTICS_URL string = EXECUTIONS_URL + ":id/aggstatistic"

	NODES_URL string = base_url + "nodes/"

	BUILD_PROJECT_URL string = base_url + "build-project/"

	PROJECTS_URL string = base_url + "projects/"

	EDIT_PROJECT_URL string = PROJECTS_URL + ":key/edit"

)
