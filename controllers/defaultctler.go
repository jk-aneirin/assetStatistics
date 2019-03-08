package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "assets.example.net"
	c.Data["Email"] = "infra@example.com"
	c.TplName = "index.tpl"
	c.Render()
}
