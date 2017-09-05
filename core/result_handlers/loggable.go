package result_handlers

import (
	"bitbucket.org/instinctools/gluten/core"
	"bitbucket.org/instinctools/gluten/shared/logging"
)

type LoggableResultHandler struct {
}

func (h LoggableResultHandler) Handle(result core.StepResult) {
	logging.WithFields(logging.Fields{
		"result": result,
	}).Info("Step has been handled")
}
