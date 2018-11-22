package main

import (
	"fmt"
	"math/rand"
	"sync"
	//原子性
	"sync/atomic"
	"time"
)

var c sync.Mutex
var e sync.RWMutex

//互斥锁
func testMapSort() {
	var a map[int]int
	a = make(map[int]int, 10)

	a[9] = 0
	a[6] = 4
	a[1] = 8
	a[5] = 4
	a[3] = 6
	a[2] = 7
	a[4] = 5
	a[8] = 1
	a[7] = 2
	a[0] = 9
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			c.Lock()
			b[1] = rand.Intn(100)
			c.Unlock()
		}(a)
	}
	c.Lock()
	fmt.Println(a)
	c.Unlock()
	time.Sleep(time.Second)
}

//读写锁
func tsetLocak() {
	var a map[int]int
	a = make(map[int]int, 10)
	var count int32
	a[9] = 0
	a[6] = 4
	a[1] = 8
	a[5] = 4
	a[3] = 6
	a[2] = 7
	a[4] = 5
	a[8] = 1
	a[7] = 2
	a[0] = 9
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			e.Lock()
			b[1] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			e.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				e.RLock()
				time.Sleep(time.Millisecond)
				//fmt.Println(a)
				e.RUnlock()
				//每个协程加1
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}
	time.Sleep(time.Second)
	fmt.Println(atomic.LoadInt32(&count))
}

func main() {
	//testMapSort()
	tsetLocak()

}
