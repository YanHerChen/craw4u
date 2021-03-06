package craw4u

import (
	"net/http"
)

// Result contains data user need
type Result struct {
	// StatusCode is 200, 404 these http code
	StatusCode int
	// Err is the error that happened when trying connecting to url
	Err error
}

// Get return response from url
func Get(url string) *Result {
	// TODO: Allow parameters
	return Request("GET", url)
}

// Request handle any kind of request
func Request(method, url string) *Result {
	// send nil show that we don't need io.Reader
	req, e := http.NewRequest(method, url, nil)
	c := &http.Client{}
	r, e := c.Do(req)
	return &Result{
		StatusCode: r.StatusCode,
		Err:        e,
	}
}
