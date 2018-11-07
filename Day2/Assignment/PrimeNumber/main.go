package main

import (
	"fmt"
)

func Pmun(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var (
		m int
		n int
	)
	fmt.Scanf("%d%d", &n, &m)
	for i := n; i < m; i++ {
		if Pmun(i) == true {
			fmt.Printf("%d\n", i)
			continue
		}
	}
}
