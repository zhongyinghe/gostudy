package main

import (
	"fmt"
)

func main() {
	type (
		user struct {
			name string
			age  int
		}

		event func(string) bool
	)

	u := user{"Tom", 32}
	fmt.Println(u)

	var f event = func(s string) bool {
		println(s)
		return s != ""
	}

	f("abc")
}
