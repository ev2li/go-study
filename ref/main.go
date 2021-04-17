package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
	Score float32
}

func main() {
	//test101()
	//test102()
	var b int = 3
	test103(&b)

	var c *int = new(int)
	*c = 100
	test103(c)
}

func test101(){
	var a int = 200
	test(a)
}

func test102(){
	var a Student = Student{
		Name :	"stu01",
		Age :	18,
		Score: 	92,
	}

	test(a)
}

func test103(b interface{}){
	val := reflect.ValueOf(b)
	val.Elem().SetInt(100)
	c := val.Elem().Int()
	fmt.Printf("get value interface{} %d\n", c)
	fmt.Printf("get value interface{} %s\n", val.Elem().String())
}

func test(b interface{}){
	t := reflect.TypeOf(b)
	fmt.Println(t)

	v := reflect.ValueOf(b)
	fmt.Println(v)

	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()
	if stu, ok := iv.(Student); ok{
		fmt.Printf("%v %T\n", stu, stu)
	}
}



