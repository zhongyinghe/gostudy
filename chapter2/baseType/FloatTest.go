package main

import (
	"fmt"
)

func main() {
	var a float32 = 1.1234567899
	var b float32 = 1.12345678
	var c float32 = 1.123456781

	println(a, b, c)        //输出为科学计数法
	println(a == b, b == c) //他们的结果都为真，有效位为7位吗？
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := 5.123456789
	fmt.Printf("%f, %v\n", d, d) //结果为5.123457和5.123456789,为什么会这样呢?因为d不是32位浮点型,而是默认采用64位

	var e float32 = 8.123456789
	fmt.Printf("%f, %v\n", e, e) //为什么float32是6位小数点，而上面的a,b,c却不是
}
