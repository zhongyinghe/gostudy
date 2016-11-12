package main

func f1() {
	x := 100
	println(&x)

	x, y := 200, "abc"
	println(&x, x)
	println(y)
}

func f2() {
	x := 100
	println(&x)

	{
		x, y := 200, 300
		println(&x, x, y)
	}
}

//这个函数会编译不通过
/*func f3() {
	x := 100
	println(&x)

	x := 200
	println(&x, x)
}*/

func main() {
	f1()
	f2()
	//f3()
}
