package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	test()
}

func test(){
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}
	wc, sc, nc, oc := count(string(result))
	fmt.Printf("world count:%d\n space count :%d\n number count:%d\n others count %d\n",wc, sc, nc, oc)
}

func count(str string) (worldCount, spaceCount, numberCount, otherCount int)  {
	t := []rune(str)
	for _,v := range(t){
		switch {
		case v >= 'a' && v<= 'z':
			fallthrough
		case v >= 'A' && v<= 'Z':
			worldCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}
	}
	return
}
