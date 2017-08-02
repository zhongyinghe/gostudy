package main

import "fmt"

func main() {
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := x[:]
	fmt.Println(s1)

	s2 := x[2:5]
	fmt.Println(s2)

	s3 := x[2:5:7]
	fmt.Println(s3)

	s4 := x[4:]
	fmt.Println(s4)

	s5 := x[:4]
	fmt.Println(s5)

	s6 := x[:4:6]
	fmt.Println(s6)
}
