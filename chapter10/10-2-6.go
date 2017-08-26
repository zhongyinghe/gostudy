package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b interface{} = (*int)(nil)
	iface := (*[2]uintptr)(unsafe.Pointer(&b))
	fmt.Println(iface, iface[1] == 0)
}
