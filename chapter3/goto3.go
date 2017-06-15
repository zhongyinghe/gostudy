package main

func test() {
test:
	println("test")
	println("test exit.")
}
func main() {
	for i := 0; i < 3; i++ {
	loop:
		println(i)
	}

	goto test //label test not defined
	goto loop //goto loop jumps into block
}
