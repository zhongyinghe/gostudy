package main

func add(x, y int) {
	z := x + y
	println(z)
}

func main() {
	for i := 0; i < 10; i++ {
		go add(i, i)
	}

	println("main exit!")
}
