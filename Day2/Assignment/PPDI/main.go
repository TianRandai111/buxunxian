package main

import (
	"fmt"
)

func Ppdi(n int, m int) {
	var (
		hd int
		tp int
		dp int
	)
	for i := n; i <= m; i++ {
		hd = i / 100
		tp = ((i % 100) / 10)
		dp = (i % 100) % 10
		sum := hd*hd*hd + tp*tp*tp + dp*tp*tp
		if sum == i {
			fmt.Println(i)
		}
	}
}

func main() {
	var n, m int
	fmt.Scanf("%d%d", &n, &m)
	Ppdi(n, m)
}
