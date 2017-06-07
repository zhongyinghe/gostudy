package main

/**
 *自定义类型和原先类型不是统一类型，不能够隐式转换，也不能够直接比较
 */
func main() {
	type data int
	var d data = 10
	var x int = d
	println(x)
	println(d == x)
}
