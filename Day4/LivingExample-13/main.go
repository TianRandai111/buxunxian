package main

import (
	"fmt"
)

func test(s []int) {
	s[0] = 10
	return
}

func main() {
	var a []int
	a = make([]int, 5)
	for i := 0; i < 5; i++ {
		a[i] = i
	}
	fmt.Println(a)
	test(a)
	fmt.Println("test.a=", a)
}
