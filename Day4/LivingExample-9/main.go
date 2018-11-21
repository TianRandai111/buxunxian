package main

import (
	"fmt"
)

func test1() {
	var a [10]int
	a[0] = 100
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for i, v := range a {
		fmt.Printf("a[%d]=%d\n", i, v)
	}
}

func test2() {
	var a [10]int
	b := a
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}

func test3(i [10]int) {
	i[0] = 100
	fmt.Println("This is test3 value:", i)
}

func test4(i *[10]int) {
	(*i)[0] = 200
	fmt.Println("This is test4 value:", *i)
}

func main() {
	var value [10]int
	var value1 [10]int
	test2()

	test3(value)
	fmt.Println("This is main value:", value)

	test4(&value1)
	fmt.Println("This is main value:", value1)
}
