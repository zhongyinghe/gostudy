package main

import "fmt"

func main() {
	var d [100]struct{}
	s := d[:]

	d[1] = struct{}{}
	s[2] = struct{}{}

	fmt.Println(s[3], len(s), cap(s))
}
