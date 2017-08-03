package main

func main() {
	var m map[string]int
	println(m["a"])
	m["a"] = 1 //因为m是nil所以不能对它进行写操作
}
