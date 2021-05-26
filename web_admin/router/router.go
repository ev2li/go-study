package router

import (
	"github.com/astaxie/beego"
	"go-study/web_admin/controller/AppController"
)

func init()  {
	beego.Router("/index", &AppController.AppController{}, "*:Index")
}
