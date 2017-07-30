package main

import "log"

func main() {
	defer func() {
		for {
			if err := recover(); err != nil {
				log.Println(err)
			} else {
				log.Fatalln("fatal")
			}
		}
	}()

	defer func() {
		panic("are you ok?")
	}()

	panic("i am ok")
}
