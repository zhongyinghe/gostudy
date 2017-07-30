package main

func test(x, y int) {
	z := 0

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()

		z = x / y
	}()

	println("x / y =", z)
}

func main() {
	test(5, 0)
}
