package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	for i := 0; i < 10000; i++ {
		path := fmt.Sprintf("./log/%d.txt", i)

		f, err := os.Open(path)

		if err != nil {
			log.Println(err)
			continue
		}

		defer f.Close()
	}
}
