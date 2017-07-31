package main

import (
	"fmt"
)

func main() {
	s := "雨痕"

	//byte遍历
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: [%c]\n", i, s[i])
	}

	//字符遍历
	for i, c := range s {
		fmt.Printf("%d: [%c]\n", i, c)
	}
}
