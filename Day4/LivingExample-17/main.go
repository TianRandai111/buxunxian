package main

import (
	"fmt"
	"sort"
)

func SclieIntSort() {
	c := []int{100, 213, 45, 5, 546, 986}
	fmt.Println(c)
	for i, v := range c {
		fmt.Println("int", i, v)
	}
	sort.Ints(c)
	fmt.Println("int.sort.c", c)

	for i, v := range c {
		fmt.Println(i, v)
	}
}

func SclieStringSort() {
	c := []string{"abc", "asdf", "e", "qwer", "eeee"}
	fmt.Println(c)
	for i, v := range c {
		fmt.Println(i, v)
	}
	sort.Strings(c)
	fmt.Println("string.sort.c", c)

	for i, v := range c {
		fmt.Println(i, v)
	}
}

func Sclieslelct() {
	c := []int{99, 55, 23, 11, 65, 75, 62}
	sort.Ints(c[:])
	index := sort.SearchInts(c[:], 99)
	fmt.Println(index)

}

func main() {
	SclieIntSort()
	fmt.Println("---------------------------------------------")
	SclieStringSort()
	fmt.Println("---------------------------------------------")
	Sclieslelct()
}
