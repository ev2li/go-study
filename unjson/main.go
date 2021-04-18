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
	test202()
}

func test201() (ret string, err error) {
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
		err = fmt.Errorf("json marshal failed, err:", err)
		return
	}
	ret = string(data)
	return
}

func test202() {
	data, err := test201()
	if err != nil {
		fmt.Println("Marshal failed", err)
		return
	}

	var user1 User
	err = json.Unmarshal([]byte(data), &user1)
	if err != nil {
		fmt.Println("Unmarshal failed", err)
		return
	}
	fmt.Println(user1)
}
