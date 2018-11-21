package main

import (
	"fmt"
)

func main() {
	var a [2][5]string = [2][5]string{{"a", "b", "c", "d", "f"}, {"1", "2", "3", "4", "5"}}
	fmt.Println(a)
	for k, v := range a {
		for i, v1 := range v {
			fmt.Printf("(%d,%d)=%s \n", k, i, v1)
		}
	}
}
