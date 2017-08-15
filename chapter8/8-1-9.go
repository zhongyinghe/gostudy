package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) //利用CPU的一个核
	exit := make(chan struct{})

	go func() {
		defer close(exit)

		go func() {
			println("b")
		}()

		for i := 0; i < 4; i++ {
			println("a:", i)

			if i == 1 {
				runtime.Gosched()
			}
		}
	}()

	<-exit
	println("main exit...")
}
