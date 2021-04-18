package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

func main() {
	//test101()
	//test102()
	//test103()
	//test104()
	test105()
}

func test101() {
	user1 := &User{
		UserName: "user1",
		NickName: "鬼鬼",
		Age:      18,
		Birthday: "2008/8/8",
		Sex:      "男",
		Email:    "aaa@qq.com",
		Phone:    "110",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func test102() {
	var age int = 18
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}
	fmt.Printf("%s", string(data))
}

func test103() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}
	fmt.Printf("%s", string(data))
}

func test104() {
	var s []map[string]interface{}
	m := make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"
	s = append(s, m)

	m1 := make(map[string]interface{})
	m1["username"] = "user2"
	m1["age"] = 19
	m1["sex"] = "flmale"
	s = append(s, m1)
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}
	fmt.Printf("%s", string(data))
}

func test105() {
	data, err := myjson()
	if err != nil {
		fmt.Println("json encode failed, error:", err)
		return
	}
	var m map[string]interface{}
	//m = make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("Unmarshal failed", err)
		return
	}
	fmt.Println(m)
}

func myjson() (ret string, err error) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Errorf("json marshal failed, err:", err)
		return
	}
	ret = string(data)
	return
}
