package main

import (
	"sync"
)

func main() {
	var m sync.Mutex
	m.Lock()
	{
		m.Lock()
		m.Unlock()
	}
	m.Unlock()
}
