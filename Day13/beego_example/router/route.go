package router

import (
	"github.com/TianRandai111/buxunxian/Day13/beego_example/controller/IndexController"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")

}
