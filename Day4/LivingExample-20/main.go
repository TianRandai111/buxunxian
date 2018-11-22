package main

import (
	"fmt"
)

func testMap() {
	var a []map[int]int
	a = make([]map[int]int, 10)

	for i := 0; i < 10; i++ {
		a[i] = make(map[int]int)
		for j := 0; j <= i; j++ {
			a[i][j] = j
		}
		fmt.Println(a[i])
	}

	fmt.Println(a[0])

}

func main() {
	testMap()
}

// if a[0] == nil {
// 	a[0] = make(map[int]int)
// }
// a[0][10] = 10
// fmt.Println(a)
