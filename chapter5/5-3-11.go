package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	s1 := append(s, 10)
	s2 := append(s1, 20, 30)

	fmt.Println(s, len(s), cap(s))
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
}
