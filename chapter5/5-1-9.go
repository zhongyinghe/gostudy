package main

func main() {
	s := "hello, world"
	println("string", &s)

	bs := []byte(s)
	println("string to []byte", &bs)

	s2 := string(bs)

	println("[]byte to string", &s2)

	rs := []rune(s)
	s3 := string(rs)

	println("string to []rune", &rs)
	println("[]rune to string", &s3)
}
