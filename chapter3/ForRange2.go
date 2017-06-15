package main

func main() {
	data := [3]string{"a", "b", "c"}

	for i := range data { //只返回所有
		println(i, data[i])
	}

	for _, s := range data { //只返回值
		println(s)
	}

	for range data { //仅迭代，不返回

	}
}
