package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	test101()
}

func test101() {
	CopyFile("target.txt", "source.txt")
	fmt.Println("copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}

	defer dst.Close()

	return io.Copy(dst, src)
}
