package main

import (
	"reflect"
	"testing"
)

type Data struct {
	X int
}

func (x *Data) Inc() {
	x.X++
}

var d = new(Data)
var v = reflect.ValueOf(d)
var m = v.MethodByName("Inc")

func call() {
	d.Inc()
}

func rcall() {
	m.Call(nil)
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call()
	}
}

func BenchmarkRcall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rcall()
	}
}
