package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	u1 := user{"Tom", 12}

	fmt.Println(u1)

	u2 := user{}
	u2.name = "Jerry"
	u2.age = 10
	fmt.Println(u2)

	var u3 user
	u3.name = "Simon"
	u3.age = 30
	fmt.Println(u3)
}
