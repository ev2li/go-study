package main

import "fmt"

func add(a int, arg ...int) int {
	var sum int = a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}
func nstr(a string, arg ...string) (res string) {
	res = a
	for i := 0; i < len(arg); i++ {
		res += arg[i]
	}
	return
}

func test() {
	sum := add(10, 12, 39, 30)
	fmt.Println(sum)

	str := nstr("hello ", "world ", "golang")
	fmt.Println(str)
}

func main() {
	test()
}
