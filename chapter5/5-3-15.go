package main

import "fmt"

func main() {
	b := make([]byte, 3)
	n := copy(b, "abcde")
	fmt.Println(n, b)

	/*c := make([]byte, 5)
	n = append(c, "abcde"...)

	fmt.Println(n, c)*/

	var bs []byte
	bs = append(bs, "abcde"...)

	fmt.Println(bs)
}
