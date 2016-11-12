package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["c"] = 2
	fmt.Println(m)
	x, ok := m["b"]
	fmt.Println(x, ok)
	delete(m, "a")
}
