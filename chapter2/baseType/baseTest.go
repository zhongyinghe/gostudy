package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, c := 100, 0144, 0x64 //8进制和16进制开头是以零开头的
	fmt.Println(a, b, c)
	fmt.Printf("0b%b, %#o, %#x\n", a, a, a) //格式化时是英文的b,o和x

	fmt.Println(math.MinInt8, math.MaxInt8) //这样能够输出某个类型的最大值和最小值
}
