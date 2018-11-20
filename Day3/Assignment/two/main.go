package main

import (
	"fmt"
)

func Yzh() {

	for i := 1; i <= 10000; i++ {
		var sum int = 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if sum == i {
			fmt.Println(sum)
		}
	}
}

func main() {
	Yzh()
}
