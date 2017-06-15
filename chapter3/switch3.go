package main

func main() {
	switch x := 5; x {
	case 5:
		println("a")
	case 6, 5:
		println("b")
	}
}
