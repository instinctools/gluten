package core

import (
	log "bitbucket.org/instinctools/gluten/shared/logging"
)

type LoggableResultHandler struct {
	Name string
}

func (h LoggableResultHandler) Handle(result StepResult) {
	log.WithFields(log.Fields{
		"result": result,
	}).Info("Step has been handled")
}
