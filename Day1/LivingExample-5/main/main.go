package main

import (
	"fmt"

	"github.com/TianRandai111/buxunxian/Day1/LivingExample-5/calc"
)

func main() {
	sum := calc.Add(100, 399)
	sum1 := calc.Sub(399, 100)
	fmt.Println("sum=", sum)
	fmt.Println("sum=", sum1)
}
