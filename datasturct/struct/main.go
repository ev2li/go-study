package main

import "fmt"

//结构体是一个值类型 Go里是没有class类型的

type Student struct {
	Name string //大写可导出，在包外可访问
	Age int
	score float32//在包外不可访问
}

func main() {
	//test01()
	test02()
}

func test01()  {
	var stu Student
	stu.Name = "zhangsan"
	stu.Age = 18
	stu.score = 80
	fmt.Println(stu)
	fmt.Printf("%p\n", &stu.Name)
	fmt.Printf("%p\n", &stu.Age)
	fmt.Printf("%p\n", &stu.score)
}

func test02()  {
	var stu *Student = &Student{
		Name:"lisi",
		Age:20,
		score: 29,
	}

	var stu2 = Student{
		Name:"wangwu",
		Age:20,
		score: 29,
	}
	fmt.Println(stu.Name)
	fmt.Println(stu2.Age)
}