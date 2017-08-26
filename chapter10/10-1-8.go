package main

import (
	"fmt"
	"reflect"
)

type A int
type B struct {
	A
}

func (A) Av()  {}
func (*A) Ap() {}
func (B) Bv()  {}
func (*B) Bp() {}
func main() {
	var b B
	t := reflect.TypeOf(&b)
	s := []reflect.Type{t, t.Elem()}
	for _, t2 := range s {
		fmt.Println(t2, ":")
		for i := 0; i < t2.NumMethod(); i++ {
			fmt.Println(" ", t2.Method(i))
		}
	}
}
