package main

import "fmt"

type attr struct {
	perm int
}

type file struct {
	name string
	attr
}

func main() {
	f := file{
		name: "test.dat",
		attr: attr{
			perm: 0755,
		},
	}

	f.perm = 0644
	fmt.Println(f)
}
