package main

import (
	"fmt"
	"os"
)

//源原子方式
func main() {
	fmt.Printf("len of args:%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%d]:%v\n", i, v)
	}
}
