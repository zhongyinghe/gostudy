package main

import (
	"errors"
	"log"
)

func check(x int) error {
	if x <= 0 {
		return errors.New("x <= 0")
	}

	return nil
}

func main() {
	x := 10

	//原始版本
	/*if err := check(x); err == nil {
		x++
		println(x)
	} else {
		log.Fatalln(err)
	}*/

	//重构版本
	if err := check(x); err != nil {
		log.Fatalln(err)
	}

	x++
	println(x)
}
