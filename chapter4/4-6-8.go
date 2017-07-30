package main

import (
	"runtime/debug"
)

func test() {
	panic("i am ok")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	test()
}
