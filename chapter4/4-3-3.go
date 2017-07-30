package main

import "errors"

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}

	return x / y, nil
}

func test() (int, error) {
	return div(8, 2)
}

func log(x int, e error) {
	println("come here!")
}
func main() {
	log(test())
}
