package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func (u user) Print() {
	fmt.Printf("%+v\n", u)
}

type Printer interface {
	Print()
}

func main() {
	var u user
	u.name = "Tom"
	u.age = 31

	var p Printer = u
	p.Print()
}
