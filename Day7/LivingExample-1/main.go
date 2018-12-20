package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "do balance error\n")
	fmt.Fprintf(os.Stderr, "This is error")
}
