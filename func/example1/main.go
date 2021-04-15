package main

import "fmt"

type add_func func(int, int) int

func operator(op add_func, a, b int) int {
	return op(a, b)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) (c int) {
	c = a - b
	return
}

func test() {
	c := add
	fmt.Println(c)

	sum := operator(c, 10, 20)
	fmt.Println(sum)

	n := sub(9, 5)
	fmt.Println(n)
}

func main() {
	test()
}
