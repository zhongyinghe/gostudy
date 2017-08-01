package main

import "fmt"

func main() {
	var d1 [3]int
	var d2 [2]int
	d1 = d2

	fmt.Println(d1)

}
