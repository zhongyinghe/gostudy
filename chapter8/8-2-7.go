package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			println(id, ": ready.")
			<-ready
			println(id, ": running...")
		}(i)
	}

	time.Sleep(time.Second)
	println("ready? Go!")
	close(ready)

	wg.Wait()

}
