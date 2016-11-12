//目的:注意类型覆盖造成的问题
package main

type color byte

const (
	black color = iota
	red
	blue
)

func test(c color) {
	println(c)
}

func main() {
	test(red)
	test(100)

	x := 2
	test(x) //这里报错，int不能赋给byte类型
}
