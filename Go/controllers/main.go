package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// 使っていない
func (c *MainController) Get() {
	c.Data["Website"] = "どらまば"
	c.Data["Email"] = "commentspace@gmail.com"
	c.TplName = "index.tpl"
}
