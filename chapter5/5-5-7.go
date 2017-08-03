package main

import "fmt"

type file struct {
	name string
}

type data struct {
	file
	name string
}

func main() {
	d := data{
		name: "data",
		file: file{"file"},
	}

	fmt.Println(d)

	d.name = "data2"
	d.file.name = "file2"

	fmt.Printf("%#v\n", d)
}
