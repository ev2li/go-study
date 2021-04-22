package main

import (
	"time"
	"fmt"
)

func main() {
	test101()
}

func test101()  {
	var ch chan int
	ch = make(chan int, 1)
	go func() {
		var i int
		for{
			select {
			case ch <- i:
			default:
				fmt.Println("channel is full")
				time.Sleep(time.Second)
			}
			i++
		}
	}()

	for {
		 v := <- ch
		 fmt.Println(v)

	}
}
