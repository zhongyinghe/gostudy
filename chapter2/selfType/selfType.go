package main

import (
	"fmt"
)

type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

func main() {
	f := read | exec
	fmt.Printf("%b\n", f)
	println(f)
	println(read)
	println(write)
	println(exec)
}
