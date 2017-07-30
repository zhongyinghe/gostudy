package main

import (
	"errors"
)

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}

	return x / y, nil
}

func main() {
	div(8, 2)
}
