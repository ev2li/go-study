package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main(){
	test()
}

func test(){
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}
	strSlice := strings.Split(string(result), "+")
	if(len(strSlice) != 2){
		fmt.Println("pls input a + b")
	}
	strNumber1 := strings.Trim(strSlice[0], " ")
	strNumber2 := strings.Trim(strSlice[1], " ")

	res := multi(strNumber1, strNumber2)
	fmt.Println(res)
}

func multi(str1, str2 string) (result string) {
	if len(str1) == 0 && len(str2) == 0 {
		return
	}
	var index1 = len(str1) - 1
	var index2 = len(str2) - 1
	var left int

	for index1 >= 0 && index2 >= 0 {
		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'
		sum := int(c1) + int(c2) + left
		if(sum >= 10){
			left = 1
		}else{
			left = 0
		}

		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%s%c", result, c3)
		index1--
		index2--
	}

	for index1 >= 0{
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		}else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%s%c", result, c3)
		index1--
	}

	for index2 >= 0{
		c1 := str1[index2] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		}else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%s%c", result, c3)
		index2--
	}

	if left == 1{
		result = fmt.Sprintf("1%s", result)
	}
	return
}