package router

import (
	"github.com/astaxie/beego"
	"go-study/beego_example/controller/IndexController"
)

func init()  {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
}
