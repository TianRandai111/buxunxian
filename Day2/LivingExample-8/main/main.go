package main

import "fmt"

// //这里是建立一个副本来交换值，所以内部的结果不会影响外部
func swap(a int, b int) {
	tmp := a
	a = b
	b = tmp
	fmt.Println("s1.a=", a)
	fmt.Println("s1.b=", b)
	return
}

func swap2(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	fmt.Println("s2.a=", a)
	fmt.Println("s2.b=", b)
	return
}

func main() {
	first := 100
	second := 200
	swap(first, second)
	fmt.Println("s1.first=", first)
	fmt.Println("s1.second=", second)
	swap2(&first, &second)
	fmt.Println("s2.first=", first)
	fmt.Println("s2.second=", second)

}
