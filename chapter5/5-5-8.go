package main

import "fmt"

type file struct {
	name string
}

type log struct {
	name string
}

type data struct {
	file
	log
}

func main() {
	d := data{}
	d.file.name = "file"
	d.log.name = "log"

	fmt.Printf("%#v\n", d)
}
