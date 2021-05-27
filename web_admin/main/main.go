package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-study/web_admin/model"
	_ "go-study/web_admin/router"
	"time"
)

func initDb() (err error) {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/logadmin")
	if err != nil {
		logs.Warn("open mysql failed,", err)
		return
	}

	model.InitDb(database)
	return
}

func initEtcd()(err error){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		logs.Warn("connect failed, err:", err)
		return
	}
	model.InitEtcd(cli)
	return
}

func main()  {

	err := initDb()
	if err != nil {
		logs.Warn("initDb failed, err:%v", err)
		return
	}

	err = initEtcd()
	if err != nil {
		logs.Warn("initEtcd failed, err:%v", err)
		return
	}
	beego.Run()
}
