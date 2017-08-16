package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a, b chan int = make(chan int, 3), make(chan int, 3)
	var c chan bool

	println(a == b)
	println(c == nil)
	fmt.Printf("%p, %d\n", a, unsafe.Sizeof(a))
}
