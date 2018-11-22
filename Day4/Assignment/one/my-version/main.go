package main

import (
	"fmt"
)

func MaoPao(a []int) {
	b := len(a)
	for j := 0; j <= b-1; j++ {
		//fmt.Println("j=", j)
		for i := 0; i < b-1-j; i++ {
			//fmt.Println("\ti=", i)
			//fmt.Println("\tb=", b)
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
			//fmt.Println(a)
		}

	}

}

func main() {
	var Mao_Pao []int
	Mao_Pao = make([]int, 6)
	// Mao_Pao[0] = 6
	// Mao_Pao[1] = 5
	// Mao_Pao[2] = 4
	// Mao_Pao[3] = 3
	// Mao_Pao[4] = 2
	// Mao_Pao[5] = 1
	for i := 0; i < 6; i++ {
		fmt.Scanf("%d\n", &Mao_Pao[i])
	}
	MaoPao(Mao_Pao)
	fmt.Println(Mao_Pao)

}
