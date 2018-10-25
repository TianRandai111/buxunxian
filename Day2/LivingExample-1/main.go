package main

import "fmt"

func test(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}

func main() {
	test(6)
}
