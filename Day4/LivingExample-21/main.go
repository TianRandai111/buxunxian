package main

import (
	"fmt"
	"sort"
)

func testMapSort() {
	var a map[int]int
	a = make(map[int]int, 10)
	a[9] = 0
	a[6] = 4
	a[1] = 8
	a[5] = 4
	a[3] = 6
	a[2] = 7
	a[4] = 5
	a[8] = 1
	a[7] = 2
	a[0] = 9

	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
	}
	fmt.Println("------------------------------")
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}

func main() {
	for index := 0; index < 5; index++ {
		testMapSort()
	}
}
