package main

func main() {
	var a, b [2]int
	println(a == b)

	c := [2]int{1, 2}
	d := [2]int{1, 2}

	println(c != d)

	var e, f [2]map[string]int
	println(e == f)
}
