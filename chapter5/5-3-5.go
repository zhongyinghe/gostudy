package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a []int
	b := []int{}
	println(a == nil, b == nil)

	fmt.Printf("a: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))
	fmt.Printf("b: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)))
}
