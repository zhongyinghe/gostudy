package main

import (
	"fmt"
	"reflect"
)

type FuncString func() string

func (f FuncString) String() string {
	return f()
}

func main() {
	var t fmt.Stringer = FuncString(func() string {
		return "hello, world"
	})

	fmt.Println(t) //输出hello, world
	fmt.Println(reflect.TypeOf(t))
}
