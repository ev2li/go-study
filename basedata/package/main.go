package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex
var rwLock sync.RWMutex
func main() {
	//test()
	test2()
}

func test()  {
	var a map[int]int
	a = make(map[int]int, 6)
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int){
			lock.Lock()
			b[1] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
	time.Sleep(time.Second)
}

func test2()  {
	var a map[int]int
	a = make(map[int]int, 6)
	var count int32
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int){
			rwLock.Lock()
			b[1] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			rwLock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int){
			for{
				rwLock.RLock()
				//fmt.Println(a)
				time.Sleep(time.Millisecond)
				rwLock.RUnlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}
