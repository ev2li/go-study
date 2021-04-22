package main

import (
	"net/http"
	"fmt"
)

func main() {
	test101()
}

func test101(){
	http.HandleFunc("/", hello)
	http.HandleFunc("/user/login", login)
	err := http.ListenAndServe("0.0.0.0:8881", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("hello hello")
	fmt.Fprintf(w, "hello")
}

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("login login")
	fmt.Fprintf(w, "login")
}
