//目的：如果中断iota的自增，该如何恢复
package main

func main() {
	const (
		a = iota
		b
		c = 100 //中断iota
		d
		e = iota //恢复iota
		f
	)
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}
