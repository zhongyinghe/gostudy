package main

func main() {
	s := []int{0, 1, 2, 3, 4}
	p := &s

	println((*p)[1]) //不支持p[1]这样的操作，而如果是数组则会支持.要达到这样的效果要(*p)[1]
	(*p)[1] += 100
	println((*p)[1])
}
