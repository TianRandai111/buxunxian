package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()
	var m map[string]int
	m["stu"] = 100
}

func calc() {
	for {
		fmt.Println("i am calc")
		time.Sleep(time.Second)
	}
}

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	for i := 0; i < 100; i++ {
		go test()
		go calc()
	}
	time.Sleep(time.Second * 10000)
}
