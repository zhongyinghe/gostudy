package main

type data struct {
	x int
}

func main() {
	d := data{100}
	var t interface{} = d
	println(t.(data).x)

	//p := &t.(data) // cannot take the address of t.(data)

	//t.(data).x = 200 // cannot take the address of t.(data)

	//改进方法
	var t2 interface{} = &d
	t2.(*data).x = 200
	println(t2.(*data).x)

	println(d.x)
}
