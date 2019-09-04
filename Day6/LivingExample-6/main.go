package main

/*
 * @Descripttion: 链表实现
 * @version: v0.1
 * @Author: 步荀仙
 * @Date: 2018-12-07 19:47:33
 * @LastEditors: 步荀仙
 * @LastEditTime: 2019-07-22 11:32:37
 */
func main() {
	var a Link
	for i := 0; i < 10; i++ {
		a.InsertHead(i)
		a.InsertTail(i)
	}
	a.Trans()
}
