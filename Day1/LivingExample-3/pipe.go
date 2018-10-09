package main

import "fmt"

func test_pipe() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	var t1 int
	t1 = <-pipe
	//将channel中的数据传递给t1变量，管道中的值变不存在channel中
	pipe <- 4
	fmt.Println(len(pipe), t1)
}

func main() {
	test_pipe()
}
