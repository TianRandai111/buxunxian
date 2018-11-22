package main

import (
	"fmt"
	"sort"
)

func testMap() {
	var a map[string]int
	var b map[int]string
	var c []int
	a = make(map[string]int)
	b = make(map[int]string)

	a["酒剑仙"] = 78
	a["步荀仙"] = 759
	a["天然呆"] = 0
	a["李逍遥"] = 23
	fmt.Println(a)

	for i, v := range a {
		b[v] = i
	}
	fmt.Println(b)

	for i, _ := range b {
		c = append(c, i)
	}

	sort.Ints(c)
	for _, v := range c {
		fmt.Println(v, b[v])
	}
}

func main() {
	testMap()
}
