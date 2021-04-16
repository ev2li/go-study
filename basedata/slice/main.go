package main

import "fmt"

type slice struct{
	ptr *[100]int
	len int
	cap int
}


func main() {
	//test()
	//test2()
	//testSlice()
	test5()
}


//slice chan array都是用make来初始化的
func test(){
	var slice []int
	var arr [5]int = [5]int{1,2,3,4,5}
	//slice = make([]int, 5)
	slice = arr[2:3] //左闭右开

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}

func modify(s slice){
	s.ptr[1] = 1000
}

func test2(){
	var s1 slice
	s1 = mymake(s1, 10)
	s1.ptr[0] = 100
	modify(s1)
	fmt.Println(s1.ptr)
}

func modify2(a []int){
	a[1] = 100
}

func testSlice(){
	var b []int = []int{1, 2, 3, 4}
	modify2(b)
	fmt.Println(b)
}
func mymake(s slice, cap int) slice {
	s.ptr = new([100]int)
	s.len = 0
	s.cap = 0
	return  s
}

func test5(){
	var a = [10]int{1,2,3,4}
	b := a[1:5]
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a[1])
}