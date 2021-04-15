package main

import (
	"fmt"
)

func swap(a *int, b *int){
	tmp := *a
	*a = *b
	*b = tmp
}

func main() {
	a := 1
	b := 2

	swap(&a, &b)
	fmt.Println(a, b)

	first := 100
	second := 200

	first, second = second, first
	fmt.Println(first, second)
}
