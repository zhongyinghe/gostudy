package main

import (
	"log"
)

type TestError struct{}

func (*TestError) Error() string {
	return "error"
}

func test(x int) (int, error) {
	if x < 0 {
		return 0, new(TestError)

	}

	return x, nil
}

func main() {
	x, err := test(100)
	if err != nil {
		log.Fatalln("err != nil")
	}

	println(x)
}
