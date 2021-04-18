package main

import (
	"bufio"
	"fmt"
	"os"
)

type student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	//test101()
	test102()
}

func test101() {
	var str = "stu01 18 89.92"
	var stu student
	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)
}

func test102() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed, err:", err)
		return
	}

	fmt.Printf("read str succ, ret:%s", str)
}
