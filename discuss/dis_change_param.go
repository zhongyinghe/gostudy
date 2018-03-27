/**
 *探讨变参的使用
 */
package main

import (
	"fmt"
	"reflect"
)

func show(args ...interface{}) {
	fmt.Println(reflect.TypeOf(args))
	fmt.Println(args)
}

func showInt(args ...int) {
	fmt.Println(args)
}

func main() {
	strs := []string{"hello", "world", "!"}
	//探讨不展开的传参

	show(strs)
	//结果是:[[hello world !]]
	//结论:args ...interface{}的变参,如果不展开，则把整个切片当元素传进去,这样会形成二维数组

	//================================================

	//show(strs...)
	//结果:cannot use strs (type []string) as type []interface {} in argument to show
	//结论:基类型不同，不能展开传入，但可以整体传入

	iss := make([]interface{}, 0, 9)
	for i := 1; i < 10; i++ {
		iss = append(iss, i)
	}

	show(iss)    //结果:[[1 2 3 4 5 6 7 8 9]]
	show(iss...) //结果:[1 2 3 4 5 6 7 8 9]

	sInt := []int{1, 5, 6, 7, 9}
	//showInt(sInt)   //结果:cannot use sInt (type []int) as type int in argument to showInt
	showInt(sInt...) //结果:[1 5 6 7 9]

	//showInt(iss) //结果:cannot use iss (type []interface {}) as type int in argument to showInt
	//showInt(iss...) //结果:cannot use iss (type []interface {}) as type []int in argument to showInt

	//终极结论:不展开,把整个变量传进去，如果赋值成功就可以;如果展开,则把传进的变量类型和变参的切片类型是否一致，若一致则成功，若不一致则失败(数组传进去都不行)
}
