package rest

const base_url string = "/api/"

var (
	EXECUTIONS_URL string = base_url + "executions/"

	RESULTS_URL string = base_url + "executions/:id/results/"

	STOP_EXECUTION_URL string = base_url + "executions/:id/stop/"

	AGGREGATE_STATISTICS_URL string = base_url + "executions/:id/aggstatistic"

	NODES_URL string = base_url + "nodes/"
)
