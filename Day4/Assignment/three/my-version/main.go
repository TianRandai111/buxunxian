package main

import (
	"fmt"
)

func ChaRu(m_slice []int) {
	b := len(m_slice)
	for i := 0; i < b-1; i++ {
		for i, v := range m_slice {
			fmt.Printf("i:%d,v%d\n", i, v)
		}
		fmt.Printf("----------分界线----------")
		if m_slice[b-1] > m_slice[i] && m_slice[b-1] < m_slice[i+1] || m_slice[b-1] == m_slice[i] {
			m_slice[b-1], m_slice[i+1] = m_slice[i+1], m_slice[b-1]
			fmt.Printf("下标i:%d,下标i对应的值:%d\n", i, m_slice[i])
		} else if m_slice[b-1] < m_slice[i] {
			m_slice[b-1], m_slice[i] = m_slice[i], m_slice[b-1]
		}
	}
	return
}

func main() {
	// Mao_Pao := [...]int{1, 3, 5, 7, 9, 11}
	var Mao_Pao []int
	Mao_Pao = make([]int, 6)
	Mao_Pao[0] = 1
	Mao_Pao[1] = 3
	Mao_Pao[2] = 5
	Mao_Pao[3] = 7
	Mao_Pao[4] = 9
	Mao_Pao[5] = 11
	var c int
	fmt.Scanf("%d", &c)
	Mao_Pao = append(Mao_Pao[:], c)
	ChaRu(Mao_Pao)
	fmt.Println(Mao_Pao)

}
