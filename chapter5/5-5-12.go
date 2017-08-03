package main

//注意:这个通道类型不懂
func main() {
	exit := make(chan struct{})

	go func() {
		println("hello, world")
		exit <- struct{}{}
	}()

	<-exit
	println("end.")
}
