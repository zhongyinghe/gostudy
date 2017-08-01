package main

import "fmt"

func main() {
	a := [2][2]int{
		{1, 2},
		{3, 4},
	}

	b := [...][2]int{
		{10, 20},
		{30, 40},
	}

	c := [...][2][3]int{
		{
			{10, 20, 30},
			{40, 50, 60},
		},
		{
			{70, 80, 90},
			{100, 110, 120},
		},
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
