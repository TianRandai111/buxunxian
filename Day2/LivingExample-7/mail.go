package main

import (
	"fmt"
)

func modify(a int) {
	a = 10
	fmt.Println("modify-a=", a)
	return
}
func modify1(a *int) {
	*a = 5
}
func main() {
	var a = 100
	var b chan int = make(chan int, 1)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("b=", &b)
	modify(a)
	fmt.Println("a=", a)
	modify1(&a)
	fmt.Println("modify1-a=", a)
}
