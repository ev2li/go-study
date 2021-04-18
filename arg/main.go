package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//test101()
	test102()
}

func test101() {
	fmt.Printf("len of args:%d\n", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, v)
	}
}

/**
支持命令行参数
*/
func test102() {
	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "", "pls input conf path")
	flag.IntVar(&logLevel, "d", 0, "pls input log level")
	flag.Parse()

	fmt.Println("path", confPath)
	fmt.Println("log loglevel", logLevel)
}
