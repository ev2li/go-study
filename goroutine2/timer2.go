package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//test501()
}

func test501(){
	for{
		select {
			case <- time.After(time.Millisecond):
				fmt.Println("after")
			}
	}
}

func test502(){
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	for i := 0; i < 16; i ++{
		go func(){
			for{
/*				select {
				case t := <- time.NewTicker(time.Second):
					fmt.Println("time out")
				}*/
			}
		}()
	}
	time.Sleep(time.Second * 100)
}
