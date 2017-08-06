package main

import "fmt"

type N int

func (n N) toString() string {
	return fmt.Sprintf("%#x", n) //%#x切换为16进制格式
}

func (N) test() {
	println("hi!")
}

func main() {
	var a N = 25
	println(a.toString())
	a.test()
}
