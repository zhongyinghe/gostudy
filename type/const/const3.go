//目的：了解常量定义的要求:左右类型要一致并且不能超出取值范围
package main

const (
	x, y int  = 99, -999
	b    byte = byte(x)
	n         = uint8(y)
)

func main() {

}
