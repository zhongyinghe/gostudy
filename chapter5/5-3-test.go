package main

import (
	"fmt"
	"reflect"
)

func main() {
	var bb []byte
	ss := make([]byte, 0, 5)
	fmt.Println(reflect.TypeOf(bb).Kind())
	fmt.Println(reflect.TypeOf(ss).Kind())
	//结论:bb是slice类型,即切片

	bs := make([]byte, 2, 3)
	copy(bs, []byte("abc"))
	fmt.Println(bs, len(bs), cap(bs))

	s := make([]string, 0, 5)
	copy(s, []string{"a", "b", "c"})
	fmt.Println(s, len(s), cap(s))
	//上面两个示例的结论:copy只能把数据copy到目标的切片len之内,与cap无关

	//探讨:当原容量超过目标容量时会怎么样
	nums := make([]int, 3)
	copy(nums, []int{1, 2, 3, 4, 5})
	fmt.Println(nums, len(nums), cap(nums))
	//结论:当原容量超过目标容量时可以正常复制，但是它只接受len长度的数据

}
