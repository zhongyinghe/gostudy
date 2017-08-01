package main

import (
	"fmt"
)

func main() {
	var a [4]int
	fmt.Println(a)

	b := [4]int{2, 5}
	fmt.Println(b)

	c := [4]int{5, 3: 10}

	fmt.Println(c)

	d := [...]int{1, 2, 3}
	fmt.Println(d)

	e := [...]int{10, 3: 100}
	fmt.Println(e)
}
