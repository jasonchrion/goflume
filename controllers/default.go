package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Data["Title"] = "Flume Manager 3.0"
	c.TplName = "index.tpl"
}

func (c *MainController) Get() {
	c.TplName = c.Ctx.Request.RequestURI[1:] + ".tpl"
}
