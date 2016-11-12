package main

func test(a, b int) {
	defer println("dispose...")
	println(a / b)
}

func main() {
	test(100, 10)
}
