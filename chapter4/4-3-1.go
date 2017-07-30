package main

//因为缺乏最后的返回值所以报错
func test(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
}

func main() {
	//println("go here")
	test(2)
}
