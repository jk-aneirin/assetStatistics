package controllers

import (
	"github.com/astaxie/beego"
)

// MainController definition.
type MainController struct {
	beego.Controller
}

// Get method.
func (c *MainController) Get() {
	c.Data["Website"] = "asset.me"
	c.Data["Email"] = "wm@gmail.com"
	c.TplName = "index.tpl"
	c.Render()
}
