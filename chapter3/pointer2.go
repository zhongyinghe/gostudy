package main

func main() {
	a := struct {
		x int
	}{}
	a.x = 100

	p := &a
	p.x += 100 //能够这样使用的原因是:它是引用类型
	println(p.x)
}
