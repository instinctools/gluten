package main

import (
	"fmt"

	"bitbucket.org/instinctools/gluten/share/persistence"
)

func main() {
	defer persistence.CloseConnection()
	metric1 := persistence.NewMetric("1key", "1value")
	metric2 := persistence.NewMetric("2key", "2Value")
	metrics := []persistence.Metric{metric1, metric2}
	result := persistence.NewExecutionResult(metrics)
	execution := persistence.NewExecution("params", result)

	fmt.Println(metrics)
	fmt.Println(result)
	fmt.Println(execution)

	fmt.Println(persistence.CreateExecution(execution))
	executionDB := persistence.GetExecution(1)
	fmt.Println(executionDB)
	
//	prst.DeleteExecution(executionDB)
	

	persistence.DeleteExecution(executionDB)

}
