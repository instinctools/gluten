package core

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

type LoggableResultHandler struct {
	Name string
}

func (h LoggableResultHandler) Handle(result StepResult) {
	log.WithFields(log.Fields{
		"result": result,
	}).Info("Step has been handled")
}
