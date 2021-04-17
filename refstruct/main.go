package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Student struct {
	Name string	`json:"name"`
	Age int	`json:"age"`
	Score float32	`json:"score"`
}

func (s Student) Set(name string, age int, score float32){
	s.Name = name
	s.Score = score
	s.Age = age
}

func (s Student) Print(){
	fmt.Println(s)
}

func main() {
	test101()
}

func test101(){
	var a Student = Student{
		Name: "zhangli",
		Age: 30,
		Score: 92.5,
	}

	result, _ := json.Marshal(a)
	fmt.Println("json result:", string(result))
	handlerStruct(&a)
}

func handlerStruct(a interface{}){
	val := reflect.ValueOf(a)
	tye := reflect.TypeOf(a)
/*	kd := val.Elem().Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}*/
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("guigui")
	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++{
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}

	tag := tye.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)
	num = val.Elem().NumMethod()
	fmt.Printf("struct has %d method\n", num)
	var params []reflect.Value
	val.Method(0).Call(params)
}
