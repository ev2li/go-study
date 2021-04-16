package main

import "fmt"

func main() {
	//test2()
	//test3()
	test5()
	test6()
}

func test2() {
	var a [10]int
	a[9] = 200
	fmt.Println(1)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for index, val := range a {
		fmt.Printf("a[%d]=%d\n", index, val)
	}
}

func test3(){
	var a [10]int
	b := a
	b[0] = 100
	fmt.Println(a[0])
}

func test5(){
	var a [5]int
	test4(&a)
	fmt.Println(a)
}

func test4(arr *[5]int){
	(*arr)[0] = 1000
}

func test6(){
	fab(20)
}

func fab(n int) {
	//var a = [...]int{}
	var a []int
	a = make([]int, n)
	a[0] = 1
	a[1] = 1

	for i := 2; i < n; i++{
		a[i] = a[i - 1] + a[i - 2]
	}
	for _,v := range a{
		fmt.Println(v)
	}
}