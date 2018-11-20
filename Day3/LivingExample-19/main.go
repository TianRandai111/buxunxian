package main

import (
	"fmt"
)

func add(a int, b int, arg ...int) (sum int) {
	sl := len(arg)
	sum = a + b
	for i := 0; i < sl; i++ {
		sum += arg[i]
	}
	return
}

func cont(a string, b string, arg ...string) (sum string) {
	sl := len(arg)
	sum = a + b
	for i := 0; i < sl; i++ {
		sum += arg[i]
	}
	return
}

func main() {
	int_sum := add(100, 200, 300, 400, 500, 600, 700)
	fmt.Printf("This is int value:%d\n", int_sum)
	str_sum := cont("100", "200", "300", "400", "500", "600", "700")
	fmt.Printf("This is int value:%s", str_sum)

}
