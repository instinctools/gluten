package main 

import (
	"fmt"
    "bitbucket.org/instinctools/gluten/share/prst"
)

func main() {
	defer prst.CloseConnection()
	metric1 := prst.NewMetric("1key", "1value")
	metric2 := prst.NewMetric("2key", "2Value")
	metrics := []prst.Metric{metric1, metric2}
	result := prst.NewExecutionResult(metrics)
	
	execution := prst.NewExecution("params", result)
	
	fmt.Println(metrics)
	fmt.Println(result)
	fmt.Println(execution)
	
	fmt.Println(prst.CreateExecution(execution))
	executionDB := prst.GetExecution(1)
	fmt.Println(executionDB)
	
//	prst.DeleteExecution(executionDB)
	

}
