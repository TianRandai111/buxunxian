/*
 * @Descripttion:
 * @version:
 * @Author: 步荀仙
 * @Date: 2018-12-07 19:29:12
 * @LastEditors: 步荀仙
 * @LastEditTime: 2019-07-22 11:47:02
 */
package main

import (
	"fmt"
)

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (p *Link) InsertHead(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}
	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}
	node.next = p.head
	p.head = node
}

func (p *Link) InsertTail(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}
	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}
	p.tail.next = node
	p.tail = node

}

func (p *Link) Trans() {
	q := p.head
	for q != nil {
		fmt.Println(q.data)
		q = q.next
	}
}
