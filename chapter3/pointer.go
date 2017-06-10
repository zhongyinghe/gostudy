package main

//要区别:内存地址和指针
//指针不能进行加减运算
func main() {
	x := 10
	var p *int = &x
	*p += 20

	println(p, *p)
}
