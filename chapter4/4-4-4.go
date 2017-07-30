package main

func test() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func main() {
	add := test()
	println(add(1, 2))
}
