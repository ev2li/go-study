package main

import "fmt"

type Carer interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct {
	Name string
}

type Test1 interface {
	Hello()
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BMW) DiDi() {
	fmt.Printf("%s is didi\n", p.Name)
}

func (p *BMW) Hello() {
	fmt.Printf("Hello, I'm %s\n", p.Name)
}

func main() {
	//test201()
	test202()
}

func test201()  {
	var a interface{} //空接口
	var b int
	var c float32
	a = b
	a = c   //什么车都是车
	fmt.Printf("type of a %T\n", a)
}

func test202()  {
	var car Carer
	var test1 Test1
	fmt.Println(car)

	var bmw BMW
	bmw.Name = "BMW"
	car = &bmw
	car.Run()

	test1 = &bmw
	test1.Hello()
}
