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
	var n N = 25
	fmt.Printf("main.n: %p, %d\n", &n, n)

	f1 := N.test
	f1(n, 5)

	f2 := (*N).test
	f2(&n, 7)

	//直接调用
	N.test(n, 8)
	(*N).test(&n, 10)
}
