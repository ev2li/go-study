
package main

import (
	"fmt"
	"unicode"
)

func main()  {
	s := "hello小沙河"
	hzc := 0
	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			hzc++;
		}
	}
	fmt.Println(hzc)
}