package main

import (
	"fmt"
	"strings"
)

func Strdispose(i string, l string, n string, j int) string {
	Reveice := strings.Replace(i, l, n, j)
	return Reveice
}

func main() {
	var (
		Value_str string
		Low_str   string
		New_str   string
		Jsq_int   int
		Reveice   string
	)
	fmt.Scanf("%s%s%s%d", &Value_str, &Low_str, &New_str, &Jsq_int)
	Reveice = Strdispose(Value_str, Low_str, New_str, Jsq_int)
	fmt.Println(Reveice)
}
