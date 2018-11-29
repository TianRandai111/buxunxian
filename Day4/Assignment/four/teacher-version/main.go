package main

import (
	"fmt"
)

func Pai_xu(a []int, left, right int) {
	if left >= right {
		return
	}

	val := a[left]
	k := left
	//确定val所在的位置

	for i := left + 1; i <= right; i++ {
		if a[i] < val {
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}

	a[k] = val
	Pai_xu(a, left, k-1)
	fmt.Println("k-1", a)
	Pai_xu(a, k+1, right)
	fmt.Println("k+1", a)
}

func main() {
	Kuai_Su := [...]int{85, 23, 45, 1, 99, 105, 3}
	Pai_xu(Kuai_Su[:], 0, len(Kuai_Su)-1)
	fmt.Println(Kuai_Su)
}
