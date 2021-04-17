package main

import "fmt"

func main() {
	test()
}

func test ()  {
	a := [...]int{17,2,9,21,49}
	choose(a[:])
	fmt.Println(a)
}

func choose(a []int)  {
	for i := 0; i< len(a); i++ {
		var min int = i
		for j := i+1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}