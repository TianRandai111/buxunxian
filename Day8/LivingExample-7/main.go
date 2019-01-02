package main

import (
	"fmt"
)

func calc(t chan int, r chan int, e chan bool) {
	for v := range t {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			r <- v
		}
	}
	fmt.Println("exit")
	e <- true
}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)
	go func() {
		for i := 0; i < 10000; i++ {
			intChan <- i
		}
		close(intChan)
	}()
	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan, exitChan)
	}

	//等待所有计算的goroutine全部退出
	go func() {
		for i := 0; i < 8; i++ {
			a := <-exitChan
			fmt.Printf(" goreoutine exit %v %v\n", i, a)
		}
		close(resultChan)
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}
