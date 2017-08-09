package main

import (
	"fmt"
)

type data int

func (d data) String() string {
	return fmt.Sprintf("data: %d", d)
}

func main() {
	var d data = 15
	var x interface{} = d

	if n, ok := x.(fmt.Stringer); ok { //判断是否实现某个接口
		fmt.Println(n)
	}

	if d2, ok := x.(data); ok { //转换为原始类型
		fmt.Println(d2)
	}

	//e := x.(error) //因为没有实现error接口的方法，所以报错
	//fmt.Println(e)
}
