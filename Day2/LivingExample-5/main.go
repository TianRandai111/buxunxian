package main

import (
	"fmt"
	"time"
)

const (
	Man    = 1
	Female = 2
)

func main() {
	second := time.Now().Unix()
	for i := 0; i <= 10; i++ {
		if second%Female == 0 {
			fmt.Println("Female")
		} else {
			fmt.Println("Man")
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
