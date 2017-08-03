package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	m := map[int]user{
		1: user{"Tom", 19},
	}

	//m[1].age += 1//这样直接访问操作时不行的，因为它是not addressable

	//改进-整个值返回
	u := m[1]
	u.age += 1
	m[1] = u
	fmt.Println(m)

	//使用指针
	m2 := map[int]*user{
		1: &user{"Jerry", 10},
	}

	m2[1].age++
	fmt.Println(*m2[1])
}
