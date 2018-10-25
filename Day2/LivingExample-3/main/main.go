package main

import (
	"fmt"
	//a是"github.com/TianRandai111/buxunxian/Day2/LivingExample-2/add"别名
	a "github.com/TianRandai111/buxunxian/Day2/LivingExample-2/add"
	"github.com/TianRandai111/buxunxian/Day2/LivingExample-3/add"
)

func main() {
	a.Test()
	fmt.Println("result:", add.Name)
	fmt.Println("result:", add.Age)
}
