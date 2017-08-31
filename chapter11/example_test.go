package main

import (
	"fmt"
	"testing"
)

func add(x, y int) int {
	return x + y
}

func ExampleAdd(t testing.T) {
	fmt.Println(add(1, 2))
	fmt.Println(add(2, 2))

	//Output:
	//3
	//4
}
