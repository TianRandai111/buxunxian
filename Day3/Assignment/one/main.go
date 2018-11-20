package main

import (
	"fmt"
)

func Cfb() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			cj := i * j
			fmt.Printf("%d * %d = %d\t", j, i, cj)
		}
		fmt.Printf("\n")
	}
}

func main() {
	Cfb()
}
