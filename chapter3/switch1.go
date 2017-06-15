package main

func main() {
	a, b, c, x := 1, 2, 3, 2
	switch x {
	case a, b:
		println("a | b")
	case c:
		println("c")
	case 4:
		println("d")
	default:
		println("z")
	}
}
