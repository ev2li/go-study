package main

import "fmt"

func main(){
	test()
}

func test(){
	var f = Adder()
	fmt.Println(f(1))
	fmt.Println(f(20))
	fmt.Println(f(300   ))
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}