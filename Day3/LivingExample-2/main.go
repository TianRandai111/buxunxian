package main

import (
	"fmt"
	"strings"
)

func Dir_File(str string) {
	var (
		Btf  bool
		Tdir string
	)
	Tdir = "/"
	Btf = strings.HasSuffix(str, Tdir)
	if Btf {
		fmt.Printf("This is full Dir:%s\n", str)
	} else {
		fmt.Printf("This isnot full Dir:%s\n", str)
		fmt.Printf("Its full from:%s%s\n", str, Tdir)
	}
}

func main() {
	var Dir_input string
	fmt.Scanf("%s", &Dir_input)
	Dir_File(Dir_input)
}
