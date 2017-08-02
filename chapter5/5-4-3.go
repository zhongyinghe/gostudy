package main

import (
	"fmt"
	//"strconv"
)

func main() {
	m := make(map[string]int)

	/*	for i := 0; i < 8; i++ {
			m["a"+strconv.Itoa(i)] = i
		}
	*/

	for i := 0; i < 8; i++ {
		m[string('a'+i)] = i
	}

	fmt.Println(m)

	for i := 0; i < 4; i++ {
		for k, v := range m {
			print(k, ":", v, " ")
		}

		println()
	}
}
