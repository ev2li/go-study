package main

import (
	"fmt"
	"go-study/refpro/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	//insts := make([]*balance.Instance, 10)
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}
	//var balancer balance.Balance
	var balanceName = "hash"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	/*	if conf == "random" {
			balancer = &balance.RandomBalance{}
		}else if conf == "robin" {
			balancer = &balance.RobinBalance{}
		}*/

	for {
		if inst, err := balance.DoBalance(balanceName, insts); err != nil {
			fmt.Println("do balance err:", err)
			continue
		} else {
			fmt.Println(inst)
			time.Sleep(time.Second)
		}
	}
}
