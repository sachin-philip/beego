package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "imsach.in"
	c.Data["Email"] = "me@imsach.in"
	c.TplNames = "index.tpl"
}
