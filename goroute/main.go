package main

import (
	"fmt"
	"time"
	"sync"
)
var (
	m = make(map[int]uint64)
	lock sync.Mutex
)

type task struct {
	n int
}

func main() {
	//test101()
	test102()
}

func test101(){
	go test()
	for {
		fmt.Println("I'm running in main")
		time.Sleep(1  * time.Second)
	}

}

func test(){
	var i int
	for {
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
	}

}

func test102(){
	for i := 0; i < 16; i++ {
		t := &task{n:i}
		go calc(t)
	}

	time.Sleep(10 * time.Second)
	for k, v := range m {
		fmt.Printf("%d! = %v\n", k, v)
	}
}

func calc(t *task){
	var sum uint64
	sum = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

//读多写少用读写锁 ，写比较多用互斥锁