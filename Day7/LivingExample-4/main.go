package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("F:\\go\\src\\github.com\\TianRandai111\\buxunxian\\Day7\\LivingExample-4\\test.txt")
	if err != nil {
		fmt.Println("read file err", err)
		return
	}
	//关闭代码
	defer file.Close()
	reader := bufio.NewReader(file)
	ftr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed error", err)
		return
	}
	fmt.Printf("read str succ, ret:%s\n", ftr)
}
