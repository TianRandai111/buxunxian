package main

import (
	"fmt"
)

func testCopy() {
	var a []int = []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 10)
	copy(b, a)
	fmt.Println(b)

	c := []int{1, 2, 3}
	c = append(c, b...)
	c = append(c, 8, 9, 0)
	fmt.Println(c)
}

func testeg() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 10)
	copy(s2, s1)
	fmt.Println(s2)

	s3 := []int{1, 2, 3}
	s3 = append(s3, s2...)
	s3 = append(s3, 4, 5, 6)
	fmt.Println(s3)
}

func main() {
	testCopy()
	fmt.Println("===================================")
	testeg()
}
