package main

func main() {
	b := 2
	x := 1 << b //报错，因为b不是无符号类型
	println(x)

}
