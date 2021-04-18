package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount     int
	NumberCount int
	SpaceCount  int
	OtherCount  int
}

func main() {
	test101()
}

func test101() {
	file, err := os.Open("/var/www/source/go/src/go-study/test.txt")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read file failed, err:%v", err)
			break
		}

		runeArray := []rune(str)
		for _, v := range runeArray {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumberCount++
			default:
				count.OtherCount++
			}
		}
	}

	fmt.Printf("char count:%d\n", count.ChCount)
	fmt.Printf("number count:%d\n", count.NumberCount)
	fmt.Printf("space count:%d\n", count.SpaceCount)
	fmt.Printf("other count:%d\n", count.OtherCount)
}
