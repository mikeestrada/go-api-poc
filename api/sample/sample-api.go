package sample

import (
	"net/http"
	"io"
	"fmt"
)

type SampleApi struct {
	Thing string
}

// e.g. http.HandleFunc("/health-check", HealthCheckHandler)
func (s *SampleApi) SampleApiHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

// TODO :ewilliams explain why this is better
func (s *SampleApi) SampleHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		// In the future we could report back on the status of our DB, or our cache
		// (e.g. Redis) by performing a simple PING, and include them in the response.
		//io.WriteString(w, `{"alive": "` + s.Thing + `"}`)
		fmt.Fprint(w, `{"alive": "` + s.Thing + `"}`);
	}
}
