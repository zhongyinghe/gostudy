package main

func main() {
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := x[2:5]

	for i := 0; i < len(s); i++ {
		println(s[i])
	}
}
