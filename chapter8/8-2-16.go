package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	c := make(chan int)

	go func() {
		defer close(done)

		for {
			select {
			case x, ok := <-c:
				if !ok {
					return
				}
				fmt.Println("data:", x)
			default:
				fmt.Println(time.Now())
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 5)

	c <- 100
	/*c <- 200
	c <- 300*/
	close(c)
	<-done
}
