package main

import (
	"fmt"
	"strings"
)

func Dspace(Space string) string {
	var retvalue string
	retvalue = strings.TrimSpace(Space)
	return retvalue
}

func Dht(Ht string, Nht string) string {
	reht := strings.Trim(Ht, Nht)
	return reht
}

func Dhead(h string, Nh string) string {
	reh := strings.TrimLeft(h, Nh)
	return reh
}

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
	fmt.Printf("去空格前原值:%s\n", Lsstr)
	fmt.Printf("去空格值:%s\n", Rspace)
	fmt.Printf("修改前前原值:%s\n", Jsz_str)
	fmt.Printf("去首尾字符:%s\n", Rht_str)
	fmt.Printf("去首字符:%s\n", Rhaed_str)
	fmt.Printf("去尾字符:%s\n", Rtail_str)
}
