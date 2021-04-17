package main

import (
	"fmt"
	"time"
)

type Cart struct {
	name	string
	age 	int
}

type Train struct {
	Cart
	int //只能有一个同名的匿名字段
	start time.Time
	age int
}

func main() {
	test01()
}

func test01()  {
	var t Train
	t.name = "zhangli"
	t.age = 30
	t.int = 200
	fmt.Println(t)
}
