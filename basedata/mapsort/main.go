package main

import (
	"fmt"
	"sort"
)

func main() {
	//test()
	test2()
}

func test()  {
	var a map[int]int
	a = make(map[int]int, 6)
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	var keys []int
	for k, _ := range a {
		keys = append(keys,k)
		//fmt.Println(k, v)
	}
	sort.Ints(keys)

	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}

func test2()  {
	var a map[string]int
	var b map[int]string

	a = make(map[string]int, 5)
	b = make(map[int]string, 5)
	a["abc"] = 5
	a["efs"] = 10

	for k, v := range a {
		b[v] = k
	}
	fmt.Println(b)
}
