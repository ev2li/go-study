package main

import "fmt"

func main() {
	test()
}

func test ()  {
	a := [...]int{17,2,9,21,49,25}
	squick(a[:])
	fmt.Println(a)
}

func squick(a []int)  {
	for i := 1; i< len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] > a[j-1] {
				break
			}else{
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
