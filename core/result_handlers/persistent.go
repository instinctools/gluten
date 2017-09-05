package result_handlers

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/core/persistence"
)

type PersistentResultHandler struct {
	resultRepo persistence.ResultRepo
}

func NewPersistentResultHandler(resultRepo persistence.ResultRepo) *PersistentResultHandler {
	return &PersistentResultHandler{
		resultRepo: resultRepo}
}

func (h PersistentResultHandler) Handle(result core.StepResult) {
	h.resultRepo.Create(result)
}
