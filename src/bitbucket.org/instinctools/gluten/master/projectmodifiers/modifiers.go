package projectmodifiers

import (
	core "bitbucket.org/instinctools/gluten/core"
)

type ProjectModifier interface {
	
	Modify(p core.Project) core.Project
	
}

type SplitLoadModifier struct {
	
	loadSpliter LoadSplitter
	
}

func (m *SplitLoadModifier) Modify(p core.Project) core.Project {
	for _, tstep := range p.GetAllSteps() {
		newSteps := m.loadSpliter.Split(tstep)
		groupingStep := core.NewStep(core.COMPOSITE_STEP, tstep.GetCommon().Name, nil, newSteps)
		tstep = groupingStep
	}
	
	return p
}

