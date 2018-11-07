package main

import (
	"fmt"
	"strconv"
)

func Ppdi(str string) {
	//var sum int
	var result int
	for i := 0; i < len(str); i++ {
		//sum += int(str[i]) * int(str[i]) * int(str[i])
		num := int(str[i] - '0')
		result += (num * num * num)
	}

	//字符串转换成整数
	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("com net convert %s to int\n", str)
	}
	if number == result {
		fmt.Printf("[%d] is PpDi\n", number)
	} else {
		fmt.Printf("[%d] is not PpDi\n", number)
	}

}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	Ppdi(str)
}
