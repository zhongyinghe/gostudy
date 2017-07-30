package main

func main() {
	var s string
	println(s == "")
	println(s == nil) //invalid operation: s == nil (mismatched types string and nil)
}
