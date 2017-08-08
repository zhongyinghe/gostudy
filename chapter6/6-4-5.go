package main

type N int

func (N) value()    {}
func (*N) pointer() {}

func main() {
	var p *N

	p.pointer()
	(*N)(nil).pointer()
	(*N).pointer(nil)

	//p.value() //为什么它会错误？猜测：因为它没有对象，哪怕为nil的对象它都没有;有指针未必有对象，而有对象一定就有指针
}
