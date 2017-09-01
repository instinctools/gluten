package services

import (
	"bitbucket.org/instinctools/gluten/core"
)

type Executor struct {
	localRunner core.TestRunner,
	remoteRunner core.TestRunner
}

func New() *Executor {
	return &Executor{
		localRunner: core.NewDefaultRunner(core.LoggableResultHandler{}),
		remoteRunner: 
	}
}

func NewWithRunner(runner core.TestRunner) *Executor {
	return &Executor{
		runner: runner,
	}
}

func ExecuteLocal(test interface{}) {

}

func ExecuteRemote(test interface{}) {

}
