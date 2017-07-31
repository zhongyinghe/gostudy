package main

import "strings"
import "bytes"

func test() string {
	var s string
	for i := 0; i < 1000; i++ {
		s += "a" //用+拼接字符串时，每次都要重新分配内存
	}
	return s
}

//改进方法1
func test2() string {
	s := make([]string, 1000)

	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}

	return strings.Join(s, "")
}

//改进方法2
func test3() string {
	var b bytes.Buffer
	b.Grow(1000) //事先准备足够的内存

	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}

	return b.String()
}

func main() {
	println(test())

	println(test2())

	println(test3())
}
