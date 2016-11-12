package main

import (
	"fmt"
)

type user struct {
	name string
	age  byte
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

	fmt.Println(m)
}
