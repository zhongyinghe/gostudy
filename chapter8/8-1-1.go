package main

func main() {
	go println("hello world")

	go func(s string) {
		println(s)
	}("hello, two")
}
