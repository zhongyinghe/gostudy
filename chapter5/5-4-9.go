package main

import (
	"fmt"
	"unsafe"
)

func test(x map[string]int) {
	fmt.Printf("x: %p\n", x)
}

func main() {
	m := make(map[string]int)
	test(m)
	fmt.Printf("m: %p, %d\n", m, unsafe.Sizeof(m))

	m2 := map[string]int{}
	test(m2)
	fmt.Printf("m: %p, %d\n", m2, unsafe.Sizeof(m2))
}
