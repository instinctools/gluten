package steps

import (
	"sync"

	"bitbucket.org/instinctools/gluten/core"
)

//ParallelStep ...
type ParallelStep struct {
	CompositeStep
	threads int
}

func NewParallelStep(name string, threadsCount int, subSteps []core.Step) *ParallelStep {
	return &ParallelStep{
		CompositeStep{
			core.BaseStep{name},
			subSteps,
		},
		threadsCount,
	}
}

func (step *ParallelStep) Run() []core.StepResult {
	//TODO check the correctness and make improvements
	stepResults := []core.StepResult{}
	subStepsCount := len(step.GetSubSteps())
	if subStepsCount > 0 {
		asyncResults := make(chan []core.StepResult, subStepsCount)
		var wg sync.WaitGroup
		wg.Add(subStepsCount)

		for _, s := range step.GetSubSteps() {
			go func() {
				defer wg.Done()
				asyncResults <- s.Run()
			}()
		}
		wg.Wait()
		for res := range asyncResults {
			stepResults = append(stepResults, res...)
		}
	}

	return stepResults
}
