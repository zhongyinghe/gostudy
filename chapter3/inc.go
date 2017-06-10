package main

func main() {
	a := 1
	a++
	println(a)

	/*
		错误的表达
		++a
		if (a++) > 1 {

		}*/

	p := &a
	*p++
	println(a)
}
