package main

func main() {
	var a interface{} = nil
	var b interface{} = (*int)(nil)

	println(a == nil, b == nil)
}
