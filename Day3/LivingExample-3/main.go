package main

import (
	"fmt"
	"strings"
)

func Head_Str(str string) {
	var (
		Pd      int
		strbyte string
	)
	strbyte = "l"
	Pd = strings.Index(str, strbyte)
	if Pd == -1 {
		fmt.Printf("This value does not exist ")
	} else {
		fmt.Printf("该值首次出现的位置 P%d", Pd+1)
	}
}

func main() {
	var Head string
	fmt.Scanf("%s", &Head)
	Head_Str(Head)
}
