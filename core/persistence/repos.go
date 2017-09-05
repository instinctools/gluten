package persistence

import (
	"bitbucket.org/instinctools/gluten/core"
)

type ExecutionRepo interface {
	Create(execution *core.Execution)
	Get(limit int, offset int) []core.Execution
	GetById(id string) core.Execution
	Update(execution core.Execution)
}

type ResultRepo interface {
	Create(result core.StepResult)
	Get(limit int, offset int) []core.StepResult
	GetById(id string) core.StepResult
	Update(result core.StepResult)
}
