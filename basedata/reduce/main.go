package main

import (
	"fmt"
	"time"
)

func recusive(n int)  {
	fmt.Println("hello\n")
	time.Sleep(time.Second)
	if n > 10 {
		return
	}
	recusive(n+1)
}

func test()  {
	recusive(0)
}

func main(){
	test()
}
