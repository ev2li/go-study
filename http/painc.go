package main

import (
	"net/http"
	"io"
	"log"
)

func main() {
	test401()
}

func test401(){
	http.HandleFunc("/test1", logPanics(SimpleServer))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {

	}
}

func logPanics(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request){
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught painc: %v", request.RemoteAddr,x)
			}
		}()
		handler(writer, request)
	}
}

func SimpleServer(write http.ResponseWriter, r http.Request)  {
	io.WriteString(write, "hello world")
}