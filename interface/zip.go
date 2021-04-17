package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWriter interface {
	Read()
	Write()
}

type File struct {
	
}

func (f *File) Read()  {
	fmt.Println("read data")
}

func (f *File) Write()  {
	fmt.Println("write  data")
}

func Test03(rw ReadWriter){
	rw.Read()
	rw.Write()
}

func test04(){
	var f *File
	var b interface{}
	b = f
	if v, ok := b.(ReadWriter); ok{
		fmt.Println(v,ok)
	}
}

func main()  {
	//var f File
	//Test03(&f)
	test04()
}
