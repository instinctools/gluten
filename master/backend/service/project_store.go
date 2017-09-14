package service

import (
	"bitbucket.org/instinctools/gluten/core"
)

type ProjectStore interface {
	AddProject(p *core.Project)
	GetByName(name string) *core.Project
	GetProjectNames() []string
}

type InMemoryProjectStore struct {
	projects map[string]*core.Project
}

func NewInMemoryProjectStore() *InMemoryProjectStore {
	return &InMemoryProjectStore{
		make(map[string]*core.Project),
	}
}

func (store *InMemoryProjectStore) AddProject(p *core.Project) {
	store.projects[p.Name] = p
}

func (store *InMemoryProjectStore) GetByName(name string) *core.Project {
	return store.projects[name]
}

func (store *InMemoryProjectStore) GetProjectNames() []string {
	keys := make([]string, 0, len(store.projects))
	for k := range store.projects {
		keys = append(keys, k)
	}
	return keys
}
