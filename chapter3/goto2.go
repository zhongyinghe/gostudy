package main

func main() {
start: //label start defined and not used
	for i := 0; i < 3; i++ {
		println(i)

		if i > 1 {
			goto exit
		}
	}

exit:
	println("exit.")
}
