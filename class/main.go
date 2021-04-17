package main

import "fmt"

type integer int
type Student struct {
	Name string
	Age int
	Score int
}

func main() {
	//test01()
	test02()
}

func test01()  {
	var stu Student
	stu.init("stu", 10, 200)

	stu1 := stu.get()
	fmt.Println(stu1)
}
func test02() {
	var a integer
	a = 100
	a.print()

	a.set(200)
	a.print()
}
func (p *Student) init(name string, age, score int){
	p.Name = name
	p.Age = age
	p.Score = score
	fmt.Println(p)
}

func (p *Student) get() Student {
	return *p
}

func (p integer) print(){
	fmt.Println(p)
}

func (p *integer) set(b integer){
	*p = b
}