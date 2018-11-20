package main

import (
	"fmt"
)

func fab(n int) int {
	if n <= 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}

func main() {
	for index := 0; index < 10; index++ {
		fmt.Println(fab(index))
	}

}
