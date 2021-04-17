package main

import "fmt"

type Student struct {
	Name string
	Sex string
}

func main() {
	//test101()
	test102()
}

func test101()  {
	var a interface{}
	var b int

	a = b
	c := a.(int)
	fmt.Printf("%d %T\n", c, c)
	fmt.Printf("%d %T\n", a, a)
	Test(a)

	var stu Student
	Test(stu)
}

func test102()  {
	var  b Student = Student{
		Name: "stu1",
		Sex: "female",
	}
	just(28, 3, 1.892, "this is test", b)
}

func Test(a interface{}){
	b, ok := a.(int)
	if ok == false {
		fmt.Println("convert failed")
		return
	}
	//a += 3
	fmt.Println(b)
}

func just(items ...interface{}){
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("%d params is bool, value is %v\n", index, v)
		case int, int32, int64:
			fmt.Printf("%d params is int, value is %v\n", index, v)
		case float32, float64:
			fmt.Printf("%d params is float32, value is %v\n", index, v)
		case string:
			fmt.Printf("%d params is string, value is %v\n", index, v)
		}
	}
}
