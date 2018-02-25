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
	r, e := http.Get(url)
	return &Result{
		StatusCode: r.StatusCode,
		Err:        e,
	}
}
