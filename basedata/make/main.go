package main

import "fmt"

func main(){
	test()
}

func test(){
	s1 := new([]int)
	fmt.Println(s1)

	s2 := make([]int, 10)
	fmt.Println(s2)
	*s1 = make([]int, 5)
	(*s1)[0] = 100
	(*s1)[1] = 200
	fmt.Println(s1)
	return
}


