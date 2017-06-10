package main

import "fmt"

func main() {
	a := 1.0 << 3
	fmt.Printf("%T, %v\n", a, a)

	var s uint = 3
	b := 1.0 << s //编译过程就对b进行了类型推断，根据1.0进行推断,所以b为浮点型
	fmt.Printf("%T, %v\n", b, b)

	var c int = 1.0 << s //自动将1.0转为int32类型
	fmt.Printf("%T, %v\n", c, c)
}
