package main

import (
	"fmt"
)

func Sw1(i int) {
	switch i {
	case 0, 1, 2, 3, 4, 5, 7, 8, 9, 10:
		fmt.Println("S1.a is equal 0")
		//fallthrough
	case 19:
		fmt.Println("S1.a is equal 10")
	default:
		fmt.Println("S1.a is equal default")
	}
	return
}

func Sw2(i int) {
	switch {
	case i > 0 && i < 10:
		fmt.Printf("%d > 0 and %d < 10", i, i)
	case i > 10 && i < 20:
		fmt.Printf("%d > 10 and %d < 20", i, i)
	default:
		fmt.Println("s2.default")
	}
}
func main() {
	var a int
	fmt.Scanf("%d", &a)
	Sw1(a)
	Sw2(a)

}
