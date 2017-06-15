package main

func main() {
	switch x := 5; x {
	default:
		println(x)
	case 5:
		x += 10
		println(x)
		if x >= 15 {
			break
		}

		fallthrough
	case 6:
		x += 20
		println(x)
	}
}
