package main

import "fmt"
import "unicode/utf8"

func main() {
	s := "雨痕"
	fmt.Printf("%s\n", s)
	fmt.Printf("% x, len: %d\n", s, len(s))
	println(utf8.RuneCountInString(s))
}
