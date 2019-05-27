package IndexController

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (p *IndexController) Index() {
	p.TplName = "IndexController/index.html"
	/* 	m := make(map[string]interface{})
	   	m["code"] = 200
	   	m["Message"] = "success"
	   	p.Data["json"] = m
	   	p.ServeJSON(true) */
}
