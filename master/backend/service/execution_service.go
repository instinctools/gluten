package service

import (
	"bitbucket.org/instinctools/gluten/core"
)

var (
	runner = core.NewDefaultRunner(&core.LoggableResultHandler{})
)

func ExecuteProject(project core.Project) {
	runner.Run(project)
}
