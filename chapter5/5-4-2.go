package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	m["a"] = 30
	m["c"] = 10

	fmt.Println(m)

	if v, ok := m["d"]; ok {
		println(v)
	}

	delete(m, "d")
}
