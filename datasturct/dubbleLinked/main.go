package main

import "fmt"

type Student struct {
	Name string
	left *Student
	right *Student
}

func main() {
	test01()
}

func test01()  {
	var root *Student = new(Student)
	root.Name = "zhangli"
/*	root.left = nil
	root.right = nil
*/
	var left *Student = new(Student)
	left.Name = "zhangsan"

	root.left = left

	var right *Student = new(Student)
	right.Name = "lisi"

	root.right = right

	var left1 *Student = new(Student)
	left1.Name = "zhangsan"

	left.left = left1

	var right1 *Student = new(Student)
	right1.Name = "lisi"

	left.right = right1
	trans(root)
}

func trans(root *Student)  {
	if root == nil {
		return
	}
	fmt.Println(root) //前序遍历
	trans(root.left)
	//fmt.Println(root) //中序遍历
	trans(root.right)
	//fmt.Println(root) //后序遍历
}