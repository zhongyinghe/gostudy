package main

func main() {
	a := [2]int{}
	b := [...][2]int{
		{10, 20},
		{30, 40},
		{50, 60},
	}

	println(len(a), cap(a))
	println(len(b), cap(b))
	println(len(b[1]), cap(b[1]))
}
