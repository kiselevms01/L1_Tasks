package main

import (
	"fmt"
	"sync"
	"time"
)

type Test struct {
	mu  sync.Mutex
	inc int
}

func (t *Test) Incrementer() {
	t.mu.Lock()
	t.inc++
	t.mu.Unlock()
}

func main() {
	in := Test{}

	for i := 0; i < 100000; i++ {
		go in.Incrementer()
	}

	time.Sleep(time.Second * 2)

	fmt.Println(in.inc)
}
