 package main

import "fmt"

func main(){
	test()
}

func modify(p *int){
	fmt.Println(p)
	*p = 10000
}

func test(){
	var a int = 10
	fmt.Println(&a)

	var p *int
	 p = &a
	fmt.Println(*p)

	 var b int = 400
	 p = &b
	 *p = 5
	 fmt.Println(b)

	 modify(&b)
	fmt.Println(b)
}
