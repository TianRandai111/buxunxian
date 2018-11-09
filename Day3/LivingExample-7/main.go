//数字转字符，字符转数字
package main

import (
	"fmt"
	"strconv"
)

func int_to_str(i int) string {
	its := strconv.Itoa(i)
	return its
}

func str_to_int(s string) int {
	sti, _ := strconv.Atoi(s)
	return sti
}

func main() {
	var (
		Zzs  int
		Szz  string
		jzzs string
		jszz int
	)
	fmt.Printf("接收一个数字转成字符串，在接受一个字符串转换为数字:(ASCII码不一样)")
	fmt.Scanf("%d%s", &jszz, &jzzs)

	Szz = int_to_str(jszz)

	Zzs = str_to_int(jzzs)
	fmt.Printf("-----------------------------------------------\n")
	fmt.Printf("Scanf first value: %d\nScanf first type: %T\nThis is string to int value: %s\nThis is type: %T\n", jszz, jszz, Szz, Szz)
	fmt.Printf("-----------------------------------------------\n")
	fmt.Printf("Scanf second value: %s\nScanf second type: %T\nThis is int to string value: %d\nThis is type: %T\n", jzzs, jzzs, Zzs, Zzs)
	fmt.Printf("-----------------------------------------------\n")
}
