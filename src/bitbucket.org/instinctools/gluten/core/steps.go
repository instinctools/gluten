package core

import (
	"net/http"
	"strconv"
)

const TYPE_REQUEST string = "REQUEST"
const TYPE_COMPOSITE string = "COMPOSITE"

// Requests steps
func GetRequestStepF (t TestStep) (string, []Metric) {
	resp, err := http.Get(t.Parameters["URL"])
	if err != nil {
		return TYPE_REQUEST, []Metric{Metric{Key: "STATUS", Val: err.Error()}}
	}
	resp.Body.Close()
	return TYPE_REQUEST, []Metric{Metric{Key: "STATUS", Val: resp.Status}}
}

// Composite steps
func RepeaterStepF (t TestStep) (string, []Metric) {
	total_repeats, err := strconv.Atoi(t.Parameters["REPEATS"])
	if err != nil {
		return TYPE_COMPOSITE, []Metric{Metric{Key: "SUCCESS_REPEATS", Val: "0"}}
	}
	success_repeats := 0
	for i := 0; i < total_repeats; i++ {
		for _, step := range t.Substeps {
			step.RunF(step)
		}
		success_repeats++
	}
	return TYPE_COMPOSITE, []Metric{Metric{Key: "SUCCESS_REPEATS", Val: strconv.Itoa(success_repeats)}}
}

func ParallelStepF (t TestStep) (string, []Metric) {
	total_repeats, err := strconv.Atoi(t.Parameters["PARALLEL_REQUESTS"])
	if err != nil {
		return TYPE_COMPOSITE, []Metric{Metric{Key: "SUCCESS_REQUESTS", Val: "0"}}
	}
	success_repeats := 0
	for i := 0; i < total_repeats; i++ {
		go func () {
			for _, step := range t.Substeps {
				step.RunF(step)
			}
		}()			
		success_repeats++
	}
	return TYPE_COMPOSITE, []Metric{Metric{Key: "SUCCESS_REQUESTS", Val: strconv.Itoa(success_repeats)}}
}
