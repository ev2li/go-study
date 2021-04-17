package main

import "fmt"

func main() {
	test()
}

func test()  {
	a := [...]int{17,2,9,21,49}
	bsort(a[:])
	fmt.Println(a)
}

func bsort(a []int){
	for i := 0; i< len(a); i++ {
		for j := i+1; j < len(a) - i; j++ {
			if a[j] < a[j-1] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}
