package main

import (
	"fmt"
)

func main() {
	var (
		a int = 8
		b *int
	)
	b = &a
	fmt.Printf("这是指针b的值：%d", *b)
}
