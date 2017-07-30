package main

import "log"

func catch() {
	log.Println("catch:", recover())
}

func main() {
	defer catch()
	defer log.Println(recover())
	defer recover()

	panic("I AM OK")
}
