package main

import (
	"strconv"
	"fmt"
)

func main()  {
	a, _ := strconv.ParseInt("10010100", 2, 32)
	b, _ := strconv.ParseInt("0144", 8, 32)
	c, _ := strconv.ParseInt("65", 16, 32)
	fmt.Println(a, b, c)

	println("0b" + strconv.FormatInt(a, 2))
	println("0" + strconv.FormatInt(a, 8))
	println("0x" + strconv.FormatInt(a, 16))
}
