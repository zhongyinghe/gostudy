//目的:了解枚举类型的定义
package main

func main() {
	const (
		x = iota
		y
		z
	)
	println(x)
	println(y)
	println(z)

	const (
		_, _ = iota, iota * 10
		a, b
		c, d
	)
	println(a)
	println(b)
	println(c)
	println(d)
}
