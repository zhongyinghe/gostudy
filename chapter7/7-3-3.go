package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

func main() {
	u := user{"tomcat", 32}
	fmt.Println(reflect.TypeOf(u))
}
