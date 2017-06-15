package main

func data() []int {
	println("origin data.")
	return []int{10, 20, 30}
}

func main() {
	for i, x := range data() { //daata()函数仅调用一次
		println(i, x)
	}
}
