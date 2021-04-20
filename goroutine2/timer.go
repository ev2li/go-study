package main

import (
	"fmt"
	"time"
)

func main4() {
	test401()
}

func test401(){
	var ch, ch2 chan int
	ch = make(chan int, 10)
	ch2 = make(chan int, 10)
	go func(){
		var i int
		for{
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i * i
			time.Sleep(time.Second)
			i++
		}
	}()

	for {
		select {
		case v := <- ch:
			fmt.Println(v)
		case v:= <- ch2:
			fmt.Println(v)
		case <- time.After(time.Second): //超时控制
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
			//return
		}
	}
}
