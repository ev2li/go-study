package main

import "fmt"

func main() {
	//test01()
	//test03()
	//test04()
	test05()
}

func test01(){
	var a [5]int = [...]int{1,2,3,4,5}
	s := a[1:]

	s[1] = 100
	fmt.Printf("before: len[%d] cap[%d]\n", len(s), cap(s))
	fmt.Printf("s=%p, a[1] = %p", s, &a[1])
	fmt.Println()
	fmt.Println("before a:", a)
	s = append(s, 10)
	s = append(s, 10)
	fmt.Printf("after: len[%d] cap[%d]\n", len(s), cap(s))
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)

	s[1] = 1000
	fmt.Println("after a:", a)
	fmt.Println(s)
	fmt.Printf("s=%p, a[1] = %p", s, &a[1])
}

func test03(){
	var a []int = []int{1,2,3,4,5,6}
	b := make([]int, 1)
	copy(b, a)
	fmt.Println(b)
}

func test04(){
	//strings 底层就是byte数组
	str := "hello world"
	s1 := str[0:5]
	s2 := str[6:]
	fmt.Println(s1)
	fmt.Println(s2)

	//str[0] = 'a' strings是不可变的 如果要修改字符串的值，要强转成byte[]
}

func test05(){
	var s string = "我hello world"
	s1 := []rune(s)
	s1[0] = '0'
	fmt.Println(s1)

	str := string(s1)
	fmt.Println(str)
}