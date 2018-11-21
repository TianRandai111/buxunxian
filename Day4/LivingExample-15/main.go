package main

import "fmt"

func main() {
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[5:]
	fmt.Println(s2)
}
