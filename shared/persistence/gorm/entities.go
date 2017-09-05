package gorm

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"bitbucket.org/instinctools/gluten/core"
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

func DtoToExecution(dto *Execution) *core.Execution {
	return &core.Execution{
		ID:     dto.Id,
		Status: dto.Status,
	}
}

func (dto *Result) toStepResult() *core.StepResult {
	metrics := []core.Metric{}
	for _, element := range dto.Metrics {
		metrics = append(metrics, element.toMetric())
	}
	return &core.StepResult{
		ExecutionID: dto.ExecutionID,
		Metrics:     metrics,
		StepType:    dto.StepType,
	}
}

func (dto Metric) toMetric() core.Metric {
	return core.Metric{
		Key: dto.Key,
		Val: dto.Value,
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
