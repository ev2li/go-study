package main

import "fmt"

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (p *Link) insertHead(data interface{}){
	node := &LinkNode{
		data :	data,
		next :	nil,
	}

	if(p.tail == nil && p.head == nil){
		p.tail = node
		p.head = node
		return
	}

	node.next = p.head
	p.head = node
}

func (p *Link) insertTail(data interface{}){
	node := &LinkNode{
		data :	data,
		next :	nil,
	}

	if(p.tail == nil && p.head == nil){
		p.tail = node
		p.head = node
		return
	}

	p.tail.next = node
	p.tail = node
}


func main() {
	test101()
}

func test101()  {
	var initLink Link
	for i := 0; i < 10; i++{
		//initLink.insertHead(i)
		initLink.insertTail(fmt.Sprintf("str %d", i))
	}
	initLink.Trans()
}

func (p *Link)  Trans(){
	v := p.head
	for v != nil {
		fmt.Println(v.data)
		v = v.next
	}
}