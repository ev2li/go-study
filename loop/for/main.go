package main

import "fmt"

func test(n int) {
	for i := 0; i < n+1; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}

func test2() {
	str := "hello world,中国"
	for index, val := range str {
		if index > 2 {
			continue
		}
		if index > 3 {
			break
		}
		fmt.Printf("index[%d] val[%c] len[%d]\n", index, val, len([]byte(string(val))))
	}
}

func main() {
	//test(5)
	test2()
}
