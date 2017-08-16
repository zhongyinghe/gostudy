package main

func main() {
	var a, b chan int
	c := make(chan int, 2)

	var send chan<- int = c
	var recv <-chan int = c

	b = (chan int)(recv)
	a = (chan int)(send)
}
