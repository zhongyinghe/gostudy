package main

func main() {
	a := [...]int{1, 2}
	println(&a, &a[0], &a[1])
}
