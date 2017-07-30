package main

import "errors"
import "fmt"

func div(x, y int) (z int, err error) {
	if y == 0 {
		err = errors.New("division by zero")
		return
	}

	z = x / y
	return
}

func main() {
	x, err := div(8, 0)
	println(x)
	fmt.Printf("%v\n", err)
}
