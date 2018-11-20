package main

import (
	"fmt"
)

func Change_num(i *int, j *int, x *int) {
	*i, *j = *j, *i
	*x = *x - 3
	return
}

func main() {
	var (
		test_num1   int
		test_num2   int
		test_num3   int
		test_p_num1 *int
		test_p_num2 *int
		test_p_num3 *int
	)
	fmt.Println("请输入第一个值，和第二个值，，和第三个值以空格作为分隔符")
	fmt.Scanf("%d%d%d", &test_num1, &test_num2, &test_num3)
	test_p_num1 = &test_num1
	test_p_num2 = &test_num2
	test_p_num3 = &test_num3
	Change_num(test_p_num1, test_p_num2, test_p_num3)
	fmt.Printf("这是第一个值:%d\n这是第二个值:%d\n这是第三个值:%d\n", test_num1, test_num2, test_num3)
	fmt.Printf("这是第一个值:%d\n这是第二个值:%d\n这是第三个值:%d\n", *test_p_num1, *test_p_num2, *test_p_num3)
}
