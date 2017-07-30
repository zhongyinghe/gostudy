package main

func test(f func()) {
	f()
}

func main() {
	test(func() {
		println("hello world")
	})
}
