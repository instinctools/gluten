package projectmodifiers

import (
	core "bitbucket.org/instinctools/gluten/core"
)

type ProjectModifier interface {
	Modify(p core.Project) core.Project
}

type SplitLoadModifier struct {
}

/*
func (m *SplitLoadModifier) Modify(p core.Project) core.Project {
	for _, tstep := range p.GetSteps() {
		newSteps := m.loadSpliter.Split(tstep)
		groupingStep := core.NewStep(core.composite_step, tstep.GetCommon().Name, nil, newSteps)
		tstep = groupingStep
	}

	return p
}

*/
