package main

import (
	"fmt"
)

func testMap() {
	var a []map[int]int
	a = make([]map[int]int, 5)

	if a[0] == nil {
		a[0] = make(map[int]int)
	}
	a[0][10] = 10
	fmt.Println(a)
}

func main() {
	testMap()
}
