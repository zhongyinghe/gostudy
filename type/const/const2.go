//目的:1、检查未使用常量是否报错；2、检测不同作用域是否可以定义同名常量
package main

func main() {
	const x = 123
	println(x)

	const y = 1.23
	{
		const x = "abc"
		println(x)
	}
}
