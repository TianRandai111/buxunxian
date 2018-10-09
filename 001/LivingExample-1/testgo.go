package main

import (
	"fmt"
	"time"
)

func test_print(a int) {
	fmt.Println(a)
}

func main() {
	for i := 0; i < 100; i++ {
		go test_print(i)
	}
	time.Sleep(time.Second)
}
