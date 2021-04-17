package main

import "fmt"

type People struct {
	name	string
	age		int
}

type Test interface {
	print()
	sleep()
}

type Student struct {
	name	string
	age 	int
	score	int
}

func main() {
	test01()
}

func test01()  {
	var t Test
	var stu Student = Student{
		name:	"stu1",
		age:	20,
		score:	200,
	}
	t = &stu
	t.print()

	var people People = People{
		name:	"zhangli",
		age:	30,
	}
	people.print()
}

func (p *Student) print(){
	fmt.Println("name:", p.name)
	fmt.Println("age", p.age)
	fmt.Println("score", p.score)
}

func (p Student) sleep(){}

func (p People) print(){}

func (p People) sleep(){}