package main

import "fmt"

func modify(a int){
	a = 10
	return
}
func modify1(a *int)  {
	b := 10
	*a = b
}

func main()  {
	var a = 100
	var b = make(chan int, 1)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	modify1(&a)
	fmt.Println(a)
}
