package main

import (
	"bufio"
	"fmt"
	"os"
)

func count(str string) (worldCount, spaceCount, numberCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			worldCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Scanf("%s\n", &str)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}
	wc, sc, nc, oc := count(string(result))
	fmt.Printf("大小写字母数量：%d个\n空格数量：%d个\n数字数量：%d个\n其他字符数量：%d个\n", wc, sc, nc, oc)
}
