package main

import (
	"fmt"
)

func main() {
	var a struct { //匿名结构体
		x int    `x`
		s string `s`
	}

	var b struct {
		x int
		s string
	}

	b = a

	fmt.Println(a)
	fmt.Println(b)
}
