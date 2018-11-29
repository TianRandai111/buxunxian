package main

import (
	"fmt"
)

func Xuan_Ze(a []int) {
	b := len(a)
	for i := 0; i < b-1; i++ {
		c := i
		for j := i + 1; j < b; j++ {
			if a[j] < a[c] {
				c = j
			}
		}
		if c != i {
			a[i], a[c] = a[c], a[i]
		}
	}
}

func main() {
	var Mao_Pao []int
	Mao_Pao = make([]int, 6)
	Mao_Pao[0] = 6
	Mao_Pao[1] = 5
	Mao_Pao[2] = 3
	Mao_Pao[3] = 4
	Mao_Pao[4] = 2
	Mao_Pao[5] = 1
	// for i := 0; i < 6; i++ {
	// 	fmt.Scanf("%d\n", &Mao_Pao[i])
	// }
	Xuan_Ze(Mao_Pao)
	fmt.Println(Mao_Pao)
}
