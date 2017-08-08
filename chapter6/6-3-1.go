package main

import (
	"fmt"
	"reflect"
)

type S struct{}

type T struct {
	S
}

func (S) SVal() {
	println("is sVal(")
}

func (*S) SPtr() {
	println("is sPtr()")
}

func (T) TVal() {
	println("is tVal()")
}

func (*T) TPtr() {
	println("is tPtr()")
}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)

	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main() {
	var t T
	methodSet(t)
	fmt.Println("--------------------")
	methodSet(&t)
}
