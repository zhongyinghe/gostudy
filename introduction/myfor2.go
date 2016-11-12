package main

func main() {
	x := 0
	for x < 5 {
		println(x)
		x++
	}

	y := 4
	for {
		println(y)
		y--
		if y < 0 {
			break
		}
	}
}
