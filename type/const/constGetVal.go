//目的：不指定类型和值的常量是如何取值(根据上面有值的常量)
package main

import "fmt"

func main() {
	const (
		x uint16 = 120
		y
		s = "abc"
		z
	)
	fmt.Printf("%T, %v\n", y, y) //v小写
	fmt.Printf("%T, %v\n", z, z)
}
