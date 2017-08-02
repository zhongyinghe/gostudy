package main

import "fmt"

func main() {
	s := make([]int, 10)
	for i := 0; i < 7; i++ {
		s[i] = i
	}
	fmt.Println(s)

	s1 := s[3:8]
	fmt.Println(s1, len(s1), cap(s1))

	s2 := s1[2:4:6]
	fmt.Println(s2, len(s2), cap(s2))

	s3 := s2[:1:5]

	fmt.Println(s3, len(s3), cap(s3))
}
