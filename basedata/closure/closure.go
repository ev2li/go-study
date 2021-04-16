package main

import (
	"strings"
	"fmt"
)

func main(){
	test02()
}

func test02(){
	f1 := makeSuffix(".bmp")
	fmt.Println(f1("test"))
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}