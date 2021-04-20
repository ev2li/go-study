package main

import "fmt"

func main() {
	test301()
}

func test301()  {
	intChan := make(chan int, 1000)
	resChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

		go func(){
			for i := 0; i < 100000 ; i++ {
				intChan <- i
			}
			close(intChan)
		}()

		for i := 0; i < 8; i++{
			go calc(intChan, resChan, exitChan)
	}


	//等待所有计算的goroutine全部退出
	go func(){
		for i := 0; i < 8; i++ {
			<- exitChan
		}
		close(resChan)
	}()

	go func(){
		for v := range resChan {
			fmt.Println(v)
		}
	}()
}


func calc(taskChan chan int, resChan chan int, exitChan chan bool){
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v % i == 0 {
				break
			}
		}

		if flag {
			resChan <- v
		}
	}
	exitChan <- true
}