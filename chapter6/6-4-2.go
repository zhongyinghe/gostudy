package main

import (
	"fmt"
)

type N int

func (n N) test(x int) {
	println(x)
	fmt.Printf("test.n: %p, %d\n", &n, n)
}

func main() {
	var n N = 100
	p := &n

	n++
	f1 := n.test

	n++
	f2 := p.test

	n++
	fmt.Printf("main.n: %p, %d\n", p, n)

	f1(5)
	f2(7)
}
