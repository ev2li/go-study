package main

import "fmt"

func test() {
	i := 0
	defer fmt.Println(i)
	defer fmt.Println("second")
	i++
	return
}
func test2(a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}(a, b)
	return result
}

func test3(a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}

	return result(a, b)
}

func main() {
	test()
	fmt.Println(test2(1, 3))
	fmt.Println(test3(5, 6))
}
