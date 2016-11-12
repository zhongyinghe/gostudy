//目的:了解const类型是否能够输出其地址
package main

var x = 0x100

const y = 0x200

func main() {
	println(&x, x)
	println(&y, y) //因为它在编译时就替换为值了，所以是无法从数据段中找到地址，因为它压根不在数据段中
}
