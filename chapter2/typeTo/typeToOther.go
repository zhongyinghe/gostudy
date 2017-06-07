package main

func main() {
	/*x := 100
	var b bool = x
	if x {

	}*/

	y := 100
	p := (*int)(&y)
	println(p)
}
