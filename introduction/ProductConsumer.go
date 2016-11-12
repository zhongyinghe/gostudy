package main

func consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv: ", x)
	}

	done <- true
}

func product(data chan int) {
	for i := 0; i < 10; i++ {
		println("send:", i)
		data <- i
	}

	close(data)
}

func main() {
	done := make(chan bool)
	data := make(chan int)
	go consumer(data, done)
	go product(data)

	<-done
}
