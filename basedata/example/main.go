package main

import "time"

const (
	Man = 1
	Flmale = 2
)


func main(){
	second := time.Now().Unix()
	if second / Flmale == 0 {
		println("female")
	}else{
		println("man")
	}
	time.Sleep(time.Microsecond)
}
