package main

import "fmt"

func main1() {
	test201()
}

func test201(){
	//var ch <- chan int //只读管道
	//var ch  chan <- int //只写管道
	ch := make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)
	var total int = 0
	for _ = range exitChan{
		total++
		if total == 2 {
			close(exitChan)
			break
		}
	}
}


func send1(ch chan int, exitChan chan struct{}){
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct {}
	exitChan <- a
}

func recv1(ch chan int, exitChan chan struct{}){
	for{
		v, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct {}
	exitChan <- a
}