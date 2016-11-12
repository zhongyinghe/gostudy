package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func main() {
	a, b := 10, 2
	var c int
	var err error
	c, err = div(a, b) //变量要先定义，后使用
	fmt.Println(c, err)

	c, err = div(10, 0)
	fmt.Println(c, err)
}
