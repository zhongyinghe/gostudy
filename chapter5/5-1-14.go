package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "雨痕"
	s = string(s[0:1] + s[3:4])
	fmt.Println(s, utf8.ValidString(s))
}
