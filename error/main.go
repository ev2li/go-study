package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path    string
	op      string
	ctime   string
	message string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s op=%s message=%s", p.path, p.op, p.message)
}
func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:    filename,
			op:      "read",
			ctime:   fmt.Sprintf("%v", time.Now()),
			message: err.Error(),
		}
	}
	defer file.Close()
	return nil
}
func main() {
	test101()
}

func test101() {
	err := Open("test.log")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error:", v)
	default:

	}
	if err != nil {
		fmt.Println(err)
	}

	v, ok := err.(*PathError)
	if ok {
		fmt.Println("get path error:", v)
	}

}
