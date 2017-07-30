package main

func add(x, y int) (z int) {
	{
		z := x + y
		return z
	}

	return
}

func main() {
	r := add(4, 2)
	println(r)
}
