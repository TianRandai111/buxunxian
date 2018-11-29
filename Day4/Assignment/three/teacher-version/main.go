package main

import (
	"fmt"
)

func Xuan_Ze(a []int) {
	b := len(a)
	for i := 1; i < b; i++ {
		for i, v := range a {
			fmt.Printf("i:%d,v:%d\n", i, v)
		}
		fmt.Printf("---------------------------------\n")
		for j := i; j > 0; j-- {
			if a[j] > a[j-1] {
				break
			}
			a[j], a[j-1] = a[j-1], a[j]
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
