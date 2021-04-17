package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	next *Student
	Name string //大写可导出，在包外可访问
	Age int
	score float32//在包外不可访问
}

func main() {
	//test01()
	//test02()
	test03()
}

func test01()  {
	var head Student
	head.Name ="san"
	head.Age = 19
	head.score = 100

	var stu1 Student
	stu1.Name ="zhangsan"
	stu1.Age = 21
	stu1.score = 90
	//stu1.Next = nil

	head.next = &stu1
	var p *Student = &head
	trans(p)

	var stu2 Student
	stu2.Name ="lisi"
	stu2.Age = 20
	stu2.score = 80

	stu1.next = &stu2

	trans(p)
}


func test02()  {
	var head Student
	head.Name ="san"
	head.Age = 19
	head.score = 100

	insertTail(&head)
	trans(&head)
}

func test03()  {
	//var head *Student = &Student{}
	var head *Student = new(Student)
	stu := Student{
		Name: "zhangli",
		Age: 30,
		score: 100,
	}
	insertHead(&head)
	trans(head)
	delNode(head)
	fmt.Println()
	trans(head)
	fmt.Println()
	insertNode(head, &stu)
	trans(head)
}

func insertNode(p *Student, newNode *Student){
	for p != nil {
		if p.Name == "stu5"{
			newNode.next = p.next
			p.next = newNode
			break
		}
		p = p.next
	}
}

func delNode(p *Student){
	var prev *Student = p
	for p != nil {
		if p.Name == "stu6"{
			prev.next = p.next
			break
		}
		prev = p
		p = p.next
	}

}

func trans(p *Student){
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
	fmt.Println()
}

func insertTail(p *Student){
	var tail = p
	for i := 0; i < 10; i++ {
		/*		var stu = Student{
					Name: fmt.Sprintf("stu%d", i),
					Age : rand.Intn(100),
					score: rand.Float32() * 100,
				}

				tail.next = &stu
				tail = &stu*/

		stu := &Student{
			Name: fmt.Sprintf("stu%d", i),
			Age : rand.Intn(100),
			score: rand.Float32() * 100,
		}

		tail.next = stu
		tail = stu
	}
}

func insertHead(p **Student){
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d", i),
			Age : rand.Intn(100),
			score: rand.Float32() * 100,
		}

		stu.next = *p
		*p = &stu
	}
}
