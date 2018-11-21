package main

import (
	"fmt"
)

func test() {
	var slice []int
	var arr [7]int = [...]int{0, 1, 2, 3, 4, 5, 6}

	slice = arr[2:5]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	slice = arr[0:1]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	//包含全部数组的长度
	slice = arr[:]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	//包含下标1到结尾的下边的长度
	slice = arr[1:]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

	//从开始到下标位置为3的长度
	slice = arr[:4]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)

}

func main() {
	test()
}
