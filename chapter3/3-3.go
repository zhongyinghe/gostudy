package main

import "fmt"

func main() {
	type data struct {
		x int
		s string
	}
	var a data = data{1, "abc"}
	fmt.Println(a)

	b := data{
		1,
		"abc",
	}

	fmt.Println(b)
}
