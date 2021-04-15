package main

import (
	"fmt"
	"time"
)

func test(){
	now := time.Now()
	fmt.Println(now.Format("02/1/2006 15:04"))
	fmt.Println(now.Format("2006-1-02 15:04"))
}

func test2()  {
	time.Sleep(time.Millisecond * 100)
}

func main(){
	test()

	start := time.Now().UnixNano()
	test2()
	end := time.Now().UnixNano()

	fmt.Printf("cost:%d us\n", (end - start))
}
