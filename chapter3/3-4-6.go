package main

import (
	"errors"
	"log"
	"strconv"
)

func check(s string) error {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil || n < 0 || n > 11 || n%2 != 0 {
		return errors.New("invalid number")
	}
	return nil
}
func main() {
	s := "9"
	if err := check(s); err != nil {
		log.Fatalln(err)
	}

	println("ok")
}
