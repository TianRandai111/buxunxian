package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m     = make(map[int]uint64)
	locak sync.Mutex
)

type task struct {
	n int
}

func calc(t *task) {
	var sum uint64
	sum = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}

	locak.Lock()
	m[t.n] = sum
	locak.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		t := &task{n: i}
		go calc(t)
	}
	time.Sleep(10 * time.Second)
	locak.Lock()
	for k, v := range m {
		fmt.Printf("%d! = %v\n", k, v)
	}
	locak.Unlock()
}
