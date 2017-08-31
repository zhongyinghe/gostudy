package main

import (
	"testing"
	"time"
)

func add(x, y int) int {
	return x + y
}

func BenchmarkAdd(b *testing.B) {
	time.Sleep(time.Second)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
		if i == 1 {
			b.StopTimer()
			time.Sleep(time.Second)
			b.StartTimer()
		}
	}
}
