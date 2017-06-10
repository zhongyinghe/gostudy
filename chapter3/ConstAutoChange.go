package main

import "fmt"

//无显式声明类型的常量,在表达式中回自动转换
func main() {
	const v = 20
	var a byte = 10
	b := a + v
	fmt.Printf("%T, %v\n", b, b)

	const c float32 = 1.2
	d := c + v
	fmt.Printf("%T, %v", d, d)
}
