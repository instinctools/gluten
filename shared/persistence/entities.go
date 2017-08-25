package persistence

import (
	"time"
)

type Execution struct {
	ID         uint `gorm:"primary_key"`
	Created    int64
	Parameters string
	Result     []ExecutionResult
}

type ExecutionResult struct {
	ID          uint `gorm:"primary_key"`
	Created     int64
	Metrics     []Metric
	ExecutionID uint
}

type Metric struct {
	ExecutionResultID uint
	Key               string
	Value             string
}

func (Execution) TableName() string {
	return "execution"
}

func (ExecutionResult) TableName() string {
	return "execution_result"
}

func (Metric) TableName() string {
	return "metric"
}

func NewExecution(parameters string, results []ExecutionResult) Execution {
	return Execution{
		Created:    time.Now().UnixNano(),
		Parameters: parameters,
		Result:     results,
	}
}

func NewExecutionResult(metrics []Metric) ExecutionResult {
	return ExecutionResult{
		Created: time.Now().UnixNano(),
		Metrics: metrics,
	}
}

func NewMetric(key, value string) Metric {
	return Metric{
		Key:   key,
		Value: value,
	}
}
