package main

import "fmt"

func main() {
	var bs []byte
	bs = append(bs, "abc"...) //这个字符串后面三个点是把字符串当作切片处理

	fmt.Println(bs)
	fmt.Printf("%s\n", bs)
}
