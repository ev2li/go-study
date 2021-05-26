package IndexController

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (p *IndexController)Index() {
	p.TplName = "IndexController/layout.html"
/*	res := make(map[string]interface{})
	res["code"] = 200
	res["message"] = "success"
	p.Data["json"] = &res
  	p.ServeJSON(true)*/
}