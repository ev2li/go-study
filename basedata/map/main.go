package main

import "fmt"

func main() {
	//test()
	//test2()
	//test3()
	//test4()
	test6()
}

func test()  {
	//var map1 map[string]string = map[string]string{"key":"value",}
	map1 := make(map[string]string, 10)
	map1["abc"] = "efg"
	map1["abc"] = "efg"
	map1["abc1"] = "efg"
	fmt.Println(map1)
}

func test2()  {
	//var map1 map[string]map[string]string
	map1 := make(map[string]map[string]string, 100)
	map1["key1"] = make(map[string]string)
	map1["key1"]["key2"] = "abc"
	map1["key1"]["key3"] = "abc"
	map1["key1"]["key4"] = "abc"
	map1["key1"]["key5"] = "abc"
	fmt.Println(map1)
	val, ok := map1["key1"]
	if ok {
		fmt.Println(val)
	}else {

	}
}

func test3()  {
	//var map1 map[string]map[string]string
	a:= make(map[string]map[string]string, 100)
	modify(a)
	fmt.Println(a)
}

func modify(a map[string]map[string]string){
	/*val, ok := a["zhangsan"]
	if ok {
		val["pwd"] = "123456"
		val["nickname"] = "pangpang"
	}else{
		a["zhangsan"] = make(map[string]string)
		a["zhangsan"]["pwd"] = "123456"
		a["zhangsan"]["nickname"] = "pangpang"
	}*/
	_, ok := a["zhangsan"]
	if !ok {
		a["zhangsan"] = make(map[string]string)
	}
	a["zhangsan"]["pwd"] = "123456"
	a["zhangsan"]["nickname"] = "pangpang"
}

func test4()  {
	map1 := make(map[string]map[string]string, 100)
	map1["key1"] = make(map[string]string)
	map1["key1"]["key2"] = "abc"
	map1["key1"]["key3"] = "abc"

	map1["key2"] = make(map[string]string)
	map1["key2"]["key4"] = "abc"
	map1["key2"]["key5"] = "abc"

	trans(map1)
	delete(map1, "key1")
	fmt.Println()
	trans(map1)
}

func trans(a map[string]map[string]string)  {
	for key, value := range a {
		fmt.Println(key)
		for k, v := range value {
			fmt.Println("\t",k, v)
		}
	}
}

func test6()  {
	var a []map[int]int
	a = make([]map[int]int, 6)
	if a[0] == nil {
		a[0] = make(map[int]int, 5)
	}
	a[0][10] = 5
	fmt.Println(a)
}

