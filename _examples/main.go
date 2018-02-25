package main

import (
	"fmt"

	"github.com/YanHerChen/craw4u"
)

func main() {
	r := craw4u.Get("https://dannypsnl.github.io/")
	if r.Err != nil {
		panic(r.Err)
	}
	fmt.Println(r.StatusCode)
}
