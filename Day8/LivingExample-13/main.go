package main

import (
	"fmt"
	"time"
)

func queryDB(ch chan int) {
	//time.Sleep(time.Second)
	ch <- 10000
}

func main() {
	ch := make(chan int)
	go queryDB(ch)
	t := time.NewTicker(time.Second)
	select {
	case v := <-ch:
		fmt.Println("result", v)
	case <-t.C:
		fmt.Println("timeout")
	}
	t.Stop()
}
