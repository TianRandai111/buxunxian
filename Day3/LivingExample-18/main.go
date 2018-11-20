package main

import (
	"fmt"
)

type add_func func(int, int) int

func add(a int, b int) int {
	return a + b
}

func operator(op add_func, a int, b int) int {
	c := op(a, b)
	return c
}

func main() {
	c := add
	fmt.Println(c)
	sum := operator(c, 100, 200)
	fmt.Println(sum)
}
