/**
*_是空标识符，用来规避编译器对未使用变量和导入包的错误检查
*
 */
package main

import (
	"strconv"
)

func main() {
	x, _ := strconv.Atoi("12")
	println(x)
}
