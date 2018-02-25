package craw4u

import (
	"net/http"
)

type Result struct {
	StatusCode int
	Err        error
}

func Get(url string) *Result {
	r, e := http.Get(url)
	return &Result{
		StatusCode: r.StatusCode,
		Err:        e,
	}
}
