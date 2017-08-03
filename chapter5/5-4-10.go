package main

import "fmt"

func test() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	return m
}

func testCap() map[int]int {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i
	}

	return m
}

func main() {
	m1 := test()
	m2 := testCap()

	fmt.Println(m1)
	fmt.Println(m2)
}
