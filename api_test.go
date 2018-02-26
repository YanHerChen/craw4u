package craw4u

import (
	"testing"
)

func TestRequest(t *testing.T) {
	r := Request("GET", "https://github.com/YanHerChen/craw4u")
	if r.StatusCode != 200 {
		t.Errorf("the project's url should active, what error happened? %v", r.Err)
	}
}
