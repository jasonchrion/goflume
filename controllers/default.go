package controllers

import (
	"github.com/astaxie/beego"
)

//MainController 页面跳转
type MainController struct {
	beego.Controller
}

//Index 首页
func (c *MainController) Index() {
	c.Data["Title"] = "GoFlume 0.1"
	c.TplName = "index.tpl"
}

//Get Get页面
func (c *MainController) Get() {
	c.TplName = c.Ctx.Request.RequestURI[1:] + ".tpl"
}
