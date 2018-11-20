package main

import (
	"fmt"
	"strconv"
)

//判断字符串是否都为数字
func pd(num_str1, num_str2 string) bool {
	a := []rune(num_str1)
	b := []rune(num_str2)
	num_len1 := len(a)
	num_len2 := len(b)
	for i := 0; i < num_len1; i++ {
		if a[i] < 48 || a[i] > 57 {
			fmt.Println("您输入的参数1不是数字")
		}
	}
	for i := 0; i < num_len2; i++ {
		if b[i] < 48 || b[i] > 57 {
			fmt.Println("您输入的参数2不是数字")
			return false
		}
	}
	return true
}

//将两个字符串前+0加至长度相等
func tl(str1, str2 string) (cl_str1, cl_str2 string) {
	var jl1, jl2 string
	str1_len := len(str1)
	str2_len := len(str2)

	switch {
	case str1_len > str2_len:
		js := str1_len - str2_len
		for i := 0; i <= js; i++ {
			jl2 += "0"
		}
		cl_str2 = jl2 + str2
		cl_str1 = "0" + str1
	case str1_len < str2_len:
		js := str2_len - str1_len
		for i := 0; i <= js; i++ {
			jl1 += "0"
		}
		cl_str1 = jl1 + str1
		cl_str2 = "0" + str2
	case str1_len == str2_len:
		cl_str1 = "0" + str1
		cl_str2 = "0" + str2
	}

	return
}

//数字相加
func add(one_num, two_num string) (sum_str string) {
	var num_3 int
	num_1_len := len(one_num)
	for i := 0; i < num_1_len; i++ {
		num_1 := int(one_num[num_1_len-i-1] - '0')
		num_2 := int(two_num[num_1_len-i-1] - '0')
		sum := num_1 + num_2 + num_3
		if sum >= 10 {
			num_3 = 1
			sum = sum - 10
		} else {
			num_3 = 0
		}
		sum1 := strconv.Itoa(sum)
		sum_str += sum1
	}
	return
}

//反转结果
func fz(str string) (fz_str string) {

	fz_len := len(str)
	if str[fz_len-1] == '0' {
		fz_len -= 1
	}
	for i := 0; i < fz_len; i++ {

		fz_str += fmt.Sprintf("%c", str[fz_len-i-1])
	}
	return
}

func main() {
	var num1, num2, cl_num1, cl_num2, sum, cl_sum string
	fmt.Scanf("%s%s", &num1, &num2)
	zj := pd(num1, num2)
	if zj {
		cl_num1, cl_num2 = tl(num1, num2)
		sum = add(cl_num1, cl_num2)
		cl_sum = fz(sum)
		fmt.Printf("%s+%s=%s", num1, num2, cl_sum)
	} else {
		fmt.Printf("请输入正确的数字\n")
	}
}
