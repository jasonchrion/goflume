package controllers

import (
	"goflume/conf"

	"github.com/astaxie/beego"
)

//TouristController 控制器
type TouristController struct {
	beego.Controller
}

//Get 响应请求
func (c *TouristController) Get() {
	c.Data["tourists"] = conf.TouristConfig
	c.TplName = "tourist.tpl"
}

//Page 响应请求
func (c *TouristController) Page() {
	p := c.GetString("p")
	c.TplName = "tourist" + p + ".html"
}
