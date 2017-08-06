package main

import (
	"sync"
)

type data struct {
	sync.Mutex
	buf [1024]byte
}

func main() {
	d := data{}
	d.Lock()
	defer d.Unlock()
}
