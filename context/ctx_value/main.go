package main

import (
	"fmt"
	"golang.org/x/net/context"
)

func process(ctx context.Context){
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 123456
	}
	fmt.Printf("ret:%d\n", ret)

	s, _ := ctx.Value("session").(string)
	fmt.Printf("ret:%s\n", s)
}

func main()  {
	ctx := context.WithValue(context.Background(), "trace_id",1343333)
	ctx = context.WithValue(ctx, "session", "ddsfafsdfsa")
	process(ctx)
}