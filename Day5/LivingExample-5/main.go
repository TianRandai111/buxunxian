package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student
	right *Student
}

//后续遍历
/* func Htrange(root *Student) {
	if root == nil {
		return
	}
	trange(root.left)
	trange(root.right)
	fmt.Println(root)
} */
//前序遍历
func trange(root *Student) {
	if root == nil {
		return
	}

	trange(root.left)

	trange(root.right)
	fmt.Println(root)
}

func main() {
	var root *Student = new(Student)
	root.Name = "步荀仙"
	root.Age = 18
	root.Score = 100

	var left1 *Student = new(Student)
	left1.Name = "天然呆"
	left1.Age = 20
	left1.Score = 100
	root.left = left1

	var left2 *Student = new(Student)
	left2.Name = "一际云川"
	left2.Age = 20
	left2.Score = 200
	left1.left = left2

	var left3 *Student = new(Student)
	left3.Name = "落叶知秋"
	left3.Age = 20
	left3.Score = 200
	left2.left = left3

	var right1 *Student = new(Student)
	right1.Name = "血海浮屠"
	right1.Age = 20
	right1.Score = 100
	root.right = right1

	trange(root)
}
