package main

import "fmt"

func main() {
	var i int
	fmt.Println(i)

	j := new(int)
	fmt.Println(j)
	*j = 199
	fmt.Println(*j)
}
