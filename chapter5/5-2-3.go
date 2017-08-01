package main

import "fmt"

func main() {
	type user struct {
		name string
		age  int
	}

	d := [...]user{
		{"Tom", 22},
		{"jerry", 10},
	}

	fmt.Printf("%#v\n", d)
}
