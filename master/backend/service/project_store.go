package service

import (
	"bitbucket.org/instinctools/gluten/core"
)

var (
	projects = make(map[string]*core.Project)
)

func AddProject(p *core.Project) {
	projects[p.Name] = p
}

func GetByName(name string) *core.Project {
	return projects[name]
}
