package main

func main() {
	switch x := 5; x {
	default: //编译器不会先执行default块
		x += 100
		println(x)
	case 5:
		x += 50
		println(x)
	}
}
