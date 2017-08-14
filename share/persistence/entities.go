package persistence

import (
	"time"
)

type Execution struct {
	ID         uint `gorm:"primary_key"`
	Created    int64
	Parameters string
	Result     ExecutionResult `gorm:"ForeignKey:ResultID"`
	ResultID   uint
}

type ExecutionResult struct {
	ID      uint `gorm:"primary_key"`
	Created int64
	Metrics []Metric
}

type Metric struct {
	ExecutionResultID uint
	Key               string
	Value             string
}

func (Execution) TableName() string {
	return "executions"
}

func (ExecutionResult) TableName() string {
	return "executions_results"
}

func (Metric) TableName() string {
	return "metrics"
}

func NewExecution(parameters string, result ExecutionResult) Execution {
	return Execution{
		Created:    time.Now().UnixNano(),
		Parameters: parameters,
		Result:     result,
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
