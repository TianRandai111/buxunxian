package main

import (
	"fmt"
	"strings"
)

func Urlif(str string) {
	var prefix string
	prefix = "http://"

	http_url := strings.HasPrefix(str, prefix)
	if http_url {
		fmt.Printf("This is complete URL :%s\n", str)
	} else {
		fmt.Printf("This is not complete URL :%s\n", str)
		fmt.Printf("Lts complete for URL :%s%s\n", prefix, str)
	}

}

func main() {
	var Url string
	fmt.Scanf("%s", &Url)
	Urlif(Url)
}
