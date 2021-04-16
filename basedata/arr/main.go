package main

import (
	"fmt"
	"errors"
)

func main(){
	test()
}
func  initConfig() (err error){
	return errors.New("init config failed")
}

func test(){
	var a []int
	a = append(a, 10, 20, 234)
	fmt.Println(a)
	a = append(a, a...)
	fmt.Println(a)

	defer func(){
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := initConfig()
	if err != nil {
		panic(err)
	}
	/*b := 0
	c := 100 / b
	fmt.Println(c)*/
}