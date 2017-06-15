package main

import (
	"log"
	"strconv"
)

func main() {
	s := "9"
	n, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		log.Fatalln(err)
	} else if n < 0 || n > 10 {
		log.Fatalln("invalid number")
	}

	println(n)
}
