package main

import "fmt"

func main() {
	m := map[string][2]int{
		"a": {1, 2},
	}

	fmt.Println(m)

	am := m["a"]
	s := am[:]
	fmt.Println(s)
}
