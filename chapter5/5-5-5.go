package main

import (
	"fmt"
	"os"
)

type data struct {
	os.File
}

func main() {
	d := data{
		File: os.File{},
	}

	fmt.Println(d)
}
