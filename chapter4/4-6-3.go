package main

import "log"

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	panic("i am ok")
	println("exit123")
}
