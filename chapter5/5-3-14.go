package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[6:9]
	tmp := s[4:]
	fmt.Println(tmp)
	n := copy(tmp, s1)
	fmt.Println(n, tmp)
	fmt.Println(s)

	s2 := make([]int, 6)
	m := copy(s2, s)
	fmt.Println(m, s2)
}
