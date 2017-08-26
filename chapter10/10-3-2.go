package main

import (
	"fmt"
	"reflect"
)

type X struct{}

func (X) Format(s string, a ...interface{}) string {
	return fmt.Sprintf(s, a...)
}
func main() {
	var a X
	v := reflect.ValueOf(&a)
	m := v.MethodByName("Format")
	out := m.Call([]reflect.Value{
		reflect.ValueOf("%s = %d"), // 所有参数都须处理
		reflect.ValueOf("x"),
		reflect.ValueOf(100),
	})
	fmt.Println(out)
	out = m.CallSlice([]reflect.Value{
		reflect.ValueOf("%s = %d"),
		reflect.ValueOf([]interface{}{"x", 100}),
	})
	fmt.Println(out)
}
