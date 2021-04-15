package main

import (
	"strings"
	"fmt"
	"strconv"
)

func urlProcess(url string) {
	fmt.Println(strings.HasPrefix(url, "jiayuan"))
}

func pathProcess(path string){
	fmt.Println(strings.HasSuffix(path, "test"))
}

func first(){
	str := "hello world abc"
	fmt.Println(strings.Repeat("aaa", 3))
	fmt.Println(strings.TrimSpace(" skss "))
	fmt.Println(strings.Trim("abcdeab", "ab"))
	fmt.Println(strings.TrimLeft("abcde", "ab"))
	fmt.Println(strings.TrimRight("abcde", "ab"))

	fmt.Println(strings.Split("ab,c,de", ","))
	fmt.Println(strings.Count("hello world abc", "e"))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	result := strings.Fields(str)
	for i := 0; i< len(result); i++{
		fmt.Println(result[i])
	}

	str2 := strings.Join(result, "\t")
	fmt.Println(str2)

	str2 = strconv.Itoa(100)
	fmt.Println(str2)

	number, err := strconv.Atoi(str2)
	if err != nil{
		fmt.Println("can not convert to int", err)
	}else {
		fmt.Println(number)
	}
}

func main(){
	/*var (
		url string
		path string
	)

	fmt.Scanf("%s%s", &url, &path)
	urlProcess(url)
	pathProcess(path)*/


	first()
}
