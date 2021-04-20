package main

import "testing"

func TestSava(t *testing.T){
	stu := &Student{
		Name: "zhangli",
		Sex: "man",
		Age: 10,
	}
	err := stu.Save()
	if err != nil {
		t.Fatal("save student failed, err:",err)
	}

}


func TestLoad(t *testing.T){
	stu := &Student{
		Name: "zhangli",
		Sex: "man",
		Age: 10,
	}
	err := stu.Save()
	if err != nil {
		t.Fatal("save student failed, err:",err)
	}
	stu2 := &Student{}
	err = stu2.Load()
	if err != nil {
		t.Fatal("load student failed, err:",err)
	}

	if stu.Name != stu2.Name {
		t.Fatal("load student failed, err:", err)
	}
}

