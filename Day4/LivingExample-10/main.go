package main

import (
	"fmt"
)

//自己写的
func fbnq(v int) {
	var a, b, c int
	a = 1
	b = 1
	fmt.Println(a)
	fmt.Println(b)
	for i := 0; i <= v; i++ {
		c = a + b
		b = c
		a = c - a
		fmt.Println(c)
	}
	return
}

//别人写的
func ofbnq(n int) {
	var a []uint64
	a = make([]uint64, n)
	a[0] = 1
	a[1] = 1

	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	var value int
	fmt.Scanf("%d", &value)
	//fbnq(value)
	ofbnq(value)
}
