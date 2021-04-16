package main

import (
	"fmt"
	"sort"
)

func main() {
	//test()
	//test2()
	//test3()
	test4()
}

func test()  {
	a := [...]int{3,1,5,2,9}
	sort.Ints(a[:]) //只能用切片
	fmt.Println(a)
}

func test2()  {
	var a = [...]string{"abc", "efg", "b", "A", "eeee"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func test3()  {
	var a = [...]int{1, 8, 38, 2, 348, 484}
	index := sort.SearchInts(a[:], 2)
	fmt.Println(index)
}

func test4(){
	var a = [...]float64{2.3, 0.8, 28.2, 9393.22}
	sort.Float64s(a[:])
	fmt.Println(a)
}

