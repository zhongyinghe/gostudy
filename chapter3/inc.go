package main

//自增只能作为独立语句，不能用于表达式中
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
