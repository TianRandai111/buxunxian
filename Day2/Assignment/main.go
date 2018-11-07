package main

import (
	"fmt"
)

func Gli() {
	fmt.Println("--------------------------------------")
}

func test(a *int, b int) {
	*a = 5
	b = 5
	fmt.Printf("test.*a=[%d],test.b=[%d]\n", *a, b)
	fmt.Printf("test.a=[%d]\n", a)
	fmt.Printf("test.&a=[%d],test.&b=[%d]\n", &a, &b)
}

func main() {
	var (
		a int
		b int
	)
	a = 1
	b = 2
	Gli()
	fmt.Printf("main.a=[%d],main.b=[%d]\n", a, b)
	fmt.Printf("main.&a=[%d],main.&b=[%d]\n", &a, &b)
	Gli()
	test(&a, b)
	Gli()
	fmt.Printf("main.a2=[%d],main.b2=[%d]\n", a, b)
	fmt.Printf("main.&a2=[%d],main.&b2=[%d]\n", &a, &b)
	Gli()
	//	var x, y int
	//	fmt.Scanf("%d%d", &x, &y)
	//	fmt.Println(x, y)
	var z, w int
	fmt.Scanf("%d%d", z, w)
	fmt.Println(z, w)
}
