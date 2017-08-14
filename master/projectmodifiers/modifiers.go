package projectmodifiers

import (
	core "gluten/core"
)

type ProjectModifier interface {
	
	Modify(p core.Project) core.Project
	
}

type SplitLoadModifier struct {
	
	
}

/*
func (m *SplitLoadModifier) Modify(p core.Project) core.Project {
	for _, tstep := range p.GetAllSteps() {
		newSteps := m.loadSpliter.Split(tstep)
		groupingStep := core.NewStep(core.COMPOSITE_STEP, tstep.GetCommon().Name, nil, newSteps)
		tstep = groupingStep
	}
	
	return p
}

*/