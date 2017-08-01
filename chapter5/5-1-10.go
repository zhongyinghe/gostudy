package main

import "unsafe"

func toString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func main() {
	bs := []byte("hello, world")
	s := toString(bs)

	//printDataPointer("bs: %x\n", &bs)
}
