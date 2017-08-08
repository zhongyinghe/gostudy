package main

import "fmt"

type N int

func (n N) test(x int) {
	println(x)
	fmt.Printf("test.n: %p, %d\n", &n, n)
}

func call(m func(x int)) {
	m(6)
}

func main() {
	var n N = 100
	p := &n
	fmt.Printf("main.n: %p, %v\n", p, n)

	n++
	call(n.test)

	n++
	call(p.test)
}
