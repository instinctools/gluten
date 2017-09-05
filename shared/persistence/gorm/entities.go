package gorm

import (
	"bitbucket.org/instinctools/gluten/core"
	"fmt"
	"github.com/google/uuid"
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
		Id:          uuid.New().String(),
	}
}

func (dto *Execution) toExecution() *core.Execution {
	return &core.Execution{
		ID:     dto.Id,
		Status: dto.Status,
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
		Key:      metric.Key,
		Value:    fmt.Sprint(metric.Val),
		ResultID: resultId,
	}
}
