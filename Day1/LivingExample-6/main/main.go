package main

import (
	"fmt"

	"github.com/TianRandai111/buxunxian/Day1/LivingExample-6/goroute"
)

func main() {
	var pipe chan int
	pipe = make(chan int, 1)
	go goroute.Add(100, 300, pipe)
	sum := <-pipe
	fmt.Println("sum=", sum)
}
