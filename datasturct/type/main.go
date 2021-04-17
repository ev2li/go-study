package main

import "fmt"

type integer int
type Student struct {
	Number int
}
type stu Student
func main() {
	test01()
}

func test01()  {
	var i integer = 1000
	var j int = 100
	//j = i  会报错，是两种不同的类型 如果要赋值必须进行一次强转
	fmt.Println(i)
	fmt.Println(j)

	var a Student
	a = Student{30}
	var b stu
	a = Student(b)
	fmt.Println(a)
}

/*make 用来创建map slice channel
new 用来创建值类型*/