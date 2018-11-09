/*
1.去掉空格
2.去除首位的指定字符
3.去除首部的指定字符
4.去除尾部的指定字符
5.分割字符串中以宫格分割的字符
6.分割以指定符号为分割符的字符串
7.指定分割服后用分割服将多个字符串接在一起
*/
package main

import (
	"fmt"
	"strings"
)

//指定分割服后用分割服将多个字符串接在一起
func Cjstr() {
	qp := []string{"Tian", "Ran", "dai"}
	str := strings.Join(qp, " ")
	fmt.Printf("---------------拼接字符---------------\n")
	for i, _ := range qp {
		fmt.Printf("原改值：%s,原类型:%T\n", qp[i], qp)
	}
	fmt.Printf("串接好的字符串：%s,串接类型：%T", str, str)
}

//分割空格
func Fspace() {
	qp := []string{}
	str := "Bu xun xian"
	qp = strings.Fields(str)
	fmt.Printf("---------------以空格分割---------------\n")
	fmt.Printf("原类型：%T,原值：%s\n", str, str)
	//第一种输出方式
	// count := len(qp)
	// for i := 0; i < count; i++ {
	// 	fmt.Printf("%s\n", qp[i])
	// }

	//第二种输出方式
	for i, _ := range qp {
		fmt.Printf("修改值：%s,修改类型:%T\n", qp[i], qp)
	}
}

//分割符号
func Ffh() {
	qp := []string{}
	str := "liu,fang,kun"
	fenge := ","
	qp = strings.Split(str, fenge)
	fmt.Printf("---------------以指定字符分割---------------\n")
	fmt.Printf("原值：%s,原类型：%T,分隔符：%s\n", str, str, fenge)
	count := len(qp)
	for i := 0; i < count; i++ {
		fmt.Printf("修改值：%s,修改类型:%T\n", qp[i], qp)
	}
}

//删除收尾空格
func Dspace(Space string) string {
	var retvalue string
	retvalue = strings.TrimSpace(Space)
	return retvalue
}

//删除首位指定字符
func Dht(Ht string, Nht string) string {
	reht := strings.Trim(Ht, Nht)
	return reht
}

//删除首部指定字符
func Dhead(h string, Nh string) string {
	reh := strings.TrimLeft(h, Nh)
	return reh
}

//删除尾部指定字符
func Dtail(t string, Ht string) string {
	rht := strings.TrimRight(t, Ht)
	return rht
}

func main() {
	var (
		Lsstr     string
		Jsz_str   string
		Rspace    string
		Sf_ht     string
		Sf_tail   string
		Sf_head   string
		Rht_str   string
		Rhaed_str string
		Rtail_str string
	)
	fmt.Printf("请输入值，去掉首位的字符，去掉首部字符，去掉尾部字符(以空格为分隔符):")
	fmt.Scanf("%s%s%s%s", &Jsz_str, &Sf_ht, &Sf_head, &Sf_tail)
	Lsstr = "    buxunxian"

	//去空格
	Rspace = Dspace(Lsstr)

	//去首位的字符
	Rht_str = Dht(Jsz_str, Sf_ht)

	//去首部字符
	Rhaed_str = Dhead(Jsz_str, Sf_head)

	//去尾部字符
	Rtail_str = Dtail(Jsz_str, Sf_tail)

	//输出字符
	fmt.Printf("---------------去掉空格的值---------------\n")
	fmt.Printf("去空格前原值:%s\n", Lsstr)
	fmt.Printf("去空格值:%s\n", Rspace)
	fmt.Printf("---------------去掉指定字符的值---------------\n")
	fmt.Printf("修改前前原值:%s\n", Jsz_str)
	fmt.Printf("去首尾字符:%s\n", Rht_str)
	fmt.Printf("去首字符:%s\n", Rhaed_str)
	fmt.Printf("去尾字符:%s\n", Rtail_str)
	Fspace()
	Ffh()
	Cjstr()
}
