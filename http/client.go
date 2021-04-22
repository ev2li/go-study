package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	test201()
}

func test201()  {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get data err :", err)
		return
	}
	fmt.Println(string(data))
}
