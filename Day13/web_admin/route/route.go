package router

import (
	"github.com/TianRandai111/buxunxian/Day13/web_admin/controller/AppController"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &AppController.AppController{}, "*:Index")

}
