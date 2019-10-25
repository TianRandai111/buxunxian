/*
 * @Descripttion:
 * @version:
 * @Author: 步荀仙
 * @Date: 2019-01-17 18:20:58
 * @LastEditors: 步荀仙
 * @LastEditTime: 2019-09-20 14:43:58
 */
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var mytemplate *template.Template

func simpleServer(w http.ResponseWriter, request *http.Request) {
	fmt.Println("hello hello")
	//fmt.Fprintf(w, "hello")
	p := Person{Name: "Mary", Age: 18}
	mytemplate.Execute(w, p)

}

func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}

func initTemplat(filename string) (err error) {
	mytemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

func main() {
	//初始化模板
	initTemplat("/opt/gx/go/src/github.com/TianRandai111/buxunxian/Day10/LivingExample-5/index.html")
	http.HandleFunc("/user/info", simpleServer)
	http.Handle()
	if err := http.ListenAndServe(":8088", nil); err != nil {
		fmt.Println("http listen failed")
	}
}
