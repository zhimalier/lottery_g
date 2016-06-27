package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "zhangmi.org"
	c.Data["Email"] = "james.zhangmi@gmail.com"
	c.TplName = "index.tpl"
}
