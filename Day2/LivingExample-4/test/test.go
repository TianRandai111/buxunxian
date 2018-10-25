package test

import (
	"fmt"
)

var Name string = "This is in test package"
var Age int = 1000

func init() {
	fmt.Println("This is test")
	fmt.Println("test.Package.Name=", Name)
	fmt.Println("test.Package.age=", Age)
	Age = 10
	fmt.Println("test.Package.age=", Age)
}
