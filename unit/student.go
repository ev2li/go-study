package main

import (
	"encoding/json"
	"io/ioutil"
)

type Student struct {
	Name string
	Sex string
	Age int
}


func (p *Student) Save() (err error){
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	err = ioutil.WriteFile("test.txt", data, 0755)
	return
}

func (p *Student) Load() (err error){
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, p)
	return
}