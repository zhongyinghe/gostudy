package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("./main.go")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
}
