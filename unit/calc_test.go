package main

import "testing"

func TestAdd(t *testing.T) {
	r := add(2, 3)
	if r != 5 {
		t.Fatalf("add(2, 3) error, except:%d, actual:%d",5 ,r)
	}
	t.Logf("test add sucess!")
}

func TestSub(t *testing.T) {
	r := sub(2, 3)
	if r != -1 {
		t.Fatalf("sub(2, 3) error, except:%d, actual:%d",-1 ,r)
	}
	t.Logf("test add sucess!")
}
