package main

import (
	"reflect"
	"testing"
)

type Data struct {
	X int
}

var d = new(Data)
var v = reflect.ValueOf(d).Elem()
var f = v.FieldByName("X")

func set(x int) {
	d.X = x
}

func rset(x int) {
	f.Set(reflect.ValueOf(x))
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set(100)
	}
}

func BenchmarkRset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rset(100)
	}
}
