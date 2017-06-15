package main

func count() int {
	print("count.")
	return 3
}

func main() {
	for i, c := 0, count(); i < c; i++ { //count()执行一次
		println("a", i)
	}

	c := 0

	for c < count() { //count()执行多次
		println("b", c)
		c++
	}
}
