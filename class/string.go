package main

import "fmt"


type Car1 struct {
	weight	int
	name 	string
}

type Bike1 struct {
	Car1
	wheel int
}

type Train1 struct {
	c Car1
}

func main() {
	test301()
}

func test301()  {
	var a Bike1
	a.weight = 100
	a.name = "bike"
	a.wheel = 2

	fmt.Println(a)
	a.Run1()

	var b Train1
	b.c.weight = 2000
	b.c.name =  "train"
	b.c.Run1()
	fmt.Printf("%s", b)
}

func (car1 *Car1) Run1(){
	fmt.Println("running...")
}

func (p Train1) String() string{
	str := fmt.Sprintf("name = [%s] weight=[%d]", p.c.name, p.c.weight)
	return  str
}