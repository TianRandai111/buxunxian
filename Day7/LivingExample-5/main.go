package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type charcount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	var count charcount
	file, err := os.Open("F:/go/src/github.com/TianRandai111/buxunxian/Day7/LivingExample-5/test.txt")
	if err != nil {
		fmt.Println("read file err:", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("reaad file failed ,err %v", err)
			break
		}
		runeArr := []rune(str)
		for _, v := range runeArr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v >= 'z':
				count.ChCount++
			case v >= '0' && v <= '8':
				count.NumCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Printf("char count:%d\n", count.ChCount)
	fmt.Printf("num count:%d\n", count.NumCount)
	fmt.Printf("space count:%d\n", count.SpaceCount)
	fmt.Printf("other count:%d\n", count.OtherCount)
}
