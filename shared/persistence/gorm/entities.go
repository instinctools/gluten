package gorm

import (
	"bitbucket.org/instinctools/gluten/core"
	"fmt"
	"time"
)

type Execution struct {
	Id      string `gorm:"primary_key"`
	Status  string
	Created int64
}

func NewExecution(execution *core.Execution) *Execution {
	return &Execution{
		Id:      execution.ID,
		Status:  execution.Status,
		Created: time.Now().UnixNano(),
	}
}

type Result struct {
	Id          string `gorm:"primary_key"`
	Created     int64
	Metrics     []Metric
	ExecutionID string
	StepType    string
}

func NewExecutionResult(result core.StepResult) *Result {
	metrics := []Metric{}
	for _, element := range result.Metrics {
		metrics = append(metrics, newMetric(result.ExecutionID, element))
	}
	return &Result{
		Created:     time.Now().UnixNano(),
		Metrics:     metrics,
		ExecutionID: result.ExecutionID,
		StepType:    result.StepType,
	}
}

type Metric struct {
	ResultID string
	Key      string
	Value    string
}

func (Execution) TableName() string {
	return "execution"
}

func (Result) TableName() string {
	return "execution_result"
}

func (Metric) TableName() string {
	return "metric"
}

func newMetric(resultId string, metric core.Metric) Metric {
	return Metric{
		Key:   metric.Key,
		Value: fmt.Sprint(metric.Val),
	}
}
