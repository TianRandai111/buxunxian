package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Db() {
	var i int
	rand.Seed(time.Now().Unix())
	Sjs := rand.Intn(100)
	for true {
		fmt.Scanf("%d\n", &i)
		// fmt.Println("This is i", i)
		// fmt.Println("This is Sjs", Sjs)
		switch {
		case Sjs > i:
			fmt.Println("输入的数字小了")
			// fmt.Println("This is i", i)
		case Sjs < i:
			fmt.Println("输入的数字大了")
		case Sjs == i:
			fmt.Println("猜对了")
			goto Loop
		}
	}
Loop:
}

func main() {
	Db()
}
