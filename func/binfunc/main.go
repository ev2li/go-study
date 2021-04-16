package main

import "fmt"

func main(){
	test()
}

func test(){
	var i int
	fmt.Println(i)

	j := new(int)
	*j = 100
	fmt.Println(j)
	fmt.Println(*j)
}
