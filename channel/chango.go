package main

import (
	"time"
	"fmt"
)

func main() {
	test201()
}

func test201(){
	intChan := make(chan int, 10)
	go write(intChan)
	go read(intChan)

	time.Sleep(time.Second)
}


func read(ch chan int){
	for {
		var b int
		b = <- ch
		fmt.Println(b)
	}
}

func write(ch chan int){
	for i := 0; i < 100; i++ {
		ch <- i
	}
}