package main

import (
	"fmt"
	"strings"
)

//大写字母转换小写
func Dzx(Vla string) string {
	var Dx string
	Dx = strings.ToLower(Vla)
	return Dx
}

//小写字母转换大写
func Xzd(Vla string) string {
	var Xx string
	Xx = strings.ToUpper(Vla)
	return Xx
}

//是用新字符，替换字符串中指定的字符，并且需要指定次数
func Strdispose(i string, l string, n string, j int) string {
	Reveice := strings.Replace(i, l, n, j)
	return Reveice
}

//指定字符，在字符串出现的次数
func Js(str string, n string) int {
	var cs int
	cs = strings.Count(str, n)
	return cs
}

func main() {
	var (
		Value_str string
		Low_str   string
		New_str   string
		Jsq_int   int
		Reveice   string
		Cs_int    int
		Daxie     string
		Xiaoxie   string
	)
	fmt.Scanf("%s%s%s%d", &Value_str, &Low_str, &New_str, &Jsq_int)
	Reveice = Strdispose(Value_str, Low_str, New_str, Jsq_int)
	Cs_int = Js(Reveice, New_str)
	Daxie = Xzd(Reveice)
	Xiaoxie = Dzx(Daxie)
	fmt.Printf("lod value:[%s]\nnew value[%s]\nCs [%d]\ncapital letter value:[%s]\nlowercase value:[%s]\n", Value_str, Reveice, Cs_int, Daxie, Xiaoxie)
}
