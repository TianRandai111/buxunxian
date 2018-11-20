package main

import (
	"fmt"
)

func reversel(a string) (reverse string) {
	alen := len(a)
	for i := 0; i < alen; i++ {
		reverse = reverse + fmt.Sprintf("%c", a[alen-i-1])
	}
	return
}

func procsess(str string) bool {
	//取字符rune
	t := []rune(str)
	length := len(t)
	for i, _ := range t {
		if i == length/2 {
			break
		}
		last := length - i - 1
		if t[i] != t[last] {
			return false
		}
	}
	return true
}

func main() {
	var Js string
	var bo bool
	fmt.Scanf("%s\n", &Js)
	//fz := reversel(Js)
	bo = procsess(Js)
	if bo {
		fmt.Printf("%s 是回文", Js)
	} else {
		fmt.Printf("%s 不是回文", Js)
	}
}
