package main

func main() {
	a := [...]int{1, 2}

	p := &a
	p[1] += 10
	println(a[1])
}
