package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
	}

	r, err := redis.Int(conn.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)


	//批量set
	_, err = conn.Do("MSet", "abc", 100, "efg", 300)
	if err != nil {
		fmt.Println(err)
	}


	r, err = redis.Ints(conn.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}
}