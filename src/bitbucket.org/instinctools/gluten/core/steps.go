package core

import (
	"net/http"
	log "bitbucket.org/instinctools/gluten/shared/logging"
	"strconv"
)

// Requests steps
func GetRequestStepF (t TestStep) (StepTyper, []Metric) {
	resp, err := http.Get(t.Parameters["URL"])
	log.WithFields(log.Fields{
		"location": resp.Request.URL.String(),
		"status": resp.Status,
		"err": err,
	}).Debug("Get Request Step has been invoked")
	if err != nil {
		return t, []Metric{Metric{Key: "STATUS", Val: err.Error()}}
	}
	resp.Body.Close()
	return t, []Metric{Metric{Key: "STATUS", Val: resp.Status}}
}

// Composite steps
func RepeaterStepF (t TestStep) (StepTyper, []Metric) {
	total_repeats, err := strconv.Atoi(t.Parameters["REPEATS"])
	if err != nil {
		return t, []Metric{Metric{Key: "SUCCESS_REPEATS", Val: "0"}}
	}
	success_repeats := 0
	for i := 0; i < total_repeats; i++ {
		for _, step := range t.Substeps {
			step.RunF(step)
		}
		success_repeats++
	}
	return t, []Metric{Metric{Key: "SUCCESS_REPEATS", Val: strconv.Itoa(success_repeats)}}
}
