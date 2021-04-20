package main

import "fmt"

//go里的channel是线程安全的，多个goroutine同时访问，不需要加锁


type student struct {
	name string
}

func main() {
	//test101()
	//test102()
	//test103()
	test104()
}

func test101()  {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m := make(map[string]string, 16)
	m["stu1"] = "001"
	m["stu2"] = "002"

	mapChan <- m
}

func test102(){
	var stuChan chan student
	stuChan = make(chan student, 10)
	stu := student{name : "stu01"}
	stuChan <- stu
}

func test103(){
	var stuChan chan *student
	stuChan = make(chan *student, 10)
	stu := student{name : "stu01"}
	stuChan <- &stu
}

func test104(){
	var stuChan chan interface{}
	stuChan = make(chan interface{}, 10)
	stu := student{name : "stu01"}
	stuChan <- &stu
	stuChan <- 5

	var stu01 interface{}
	stu01 = <-stuChan
	fmt.Println(stu01)

	var stu02 *student
	stu02, ok := stu01.(*student)
	if !ok {
		fmt.Println("convert failed!")
		return
	}
	fmt.Println(stu02)
}