package main

import (
	"fmt"
	"unsafe"
)

type point struct {
	x, y int
}

type value struct {
	id   int
	name string
	data []byte
	next *value
	point
}

func main() {
	v := value{
		id:    1,
		name:  "test",
		data:  []byte{1, 2, 3, 4},
		point: point{x: 100, y: 200},
	}

	fmt.Println(v)

	s := `
		v: %p ~ %x, size: %d, align: %d

		field   address        offset   size
		------+-------------+---------+---------
		id			%p			%d			%d
		name 		%p			%d			%d
		data		%p			%d			%d
		next		%p			%d			%d
		x			%p			%d			%d
		y			%p			%d			%d
	`

	// unsafe.Alignof 获取其对齐值,疑问:对齐值有什么用?
	fmt.Printf(s,
		&v, uintptr(unsafe.Pointer(&v))+unsafe.Sizeof(v), unsafe.Sizeof(v), unsafe.Alignof(v),
		&v.id, unsafe.Offsetof(v.id), unsafe.Sizeof(v.id),
		&v.name, unsafe.Offsetof(v.name), unsafe.Sizeof(v.name),
		&v.data, unsafe.Offsetof(v.data), unsafe.Sizeof(v.data),
		&v.next, unsafe.Offsetof(v.next), unsafe.Sizeof(v.next),
		&v.x, unsafe.Offsetof(v.x), unsafe.Sizeof(v.x),
		&v.y, unsafe.Offsetof(v.y), unsafe.Sizeof(v.y))
}
