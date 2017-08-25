package steps

import (
	"sync"
)

//ParallelStep ...
type ParallelStep struct {
	CompositeStep
	threads int
}

func NewParallelStep(name string, threadsCount int, subSteps []Step) *ParallelStep {
	return &ParallelStep{
		CompositeStep{BaseStep{Name: name, SubSteps: subSteps}},
		threadsCount,
	}
}

func (step *ParallelStep) Run() []StepResult {
	//TODO check the correctness and make improvements
	stepResults := []StepResult{}
	subStepsCount := len(step.GetSubSteps())
	if subStepsCount > 0 {
		asyncResults := make(chan []StepResult, subStepsCount)
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
