package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func (u user) ToString() string {
	return fmt.Sprintf("%+v", u)
}

type manager struct {
	user
	title string
}

func main() {
	var m manager
	m.name = "Tom"
	m.age = 31
	m.title = "CEO"

	println(m.ToString())
}
