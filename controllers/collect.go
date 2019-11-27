package controllers

import (
	"goflume/models"
	"goflume/utils"

	"github.com/astaxie/beego"
)

//CollectController 控制器
type CollectController struct {
	beego.Controller
}

//Get 响应请求
func (c *CollectController) Get() {
	c.Data["collects"] = utils.LoadCollect()
	c.TplName = "collect.tpl"
}

//Update 响应请求
func (c *CollectController) Update() {
	id := c.GetString("cid")
	c.Data["c"] = utils.GetCollect(id)
	c.Data["templates"] = utils.LoadTemplate()
	c.TplName = "collectForm.tpl"
}

//New 响应请求
func (c *CollectController) New() {
	c.Data["templates"] = utils.LoadTemplate()
	c.TplName = "collectForm.tpl"
}

//Save 响应请求
func (c *CollectController) Save() {
	var ci models.CollectInfo
	c.ParseForm(&ci)
	if "" == ci.ID {
		ci.CreateTime = utils.GetTimeNow()
		ci.UpdateTime = ci.CreateTime
		ci.ID = utils.Md5("C-" + ci.Name + ci.DESC + ci.CreateTime)
	} else {
		ci.UpdateTime = utils.GetTimeNow()
	}
	utils.SaveCollect(ci)
	c.Get()
}

//Delete 响应请求
func (c *CollectController) Delete() {
	id := c.GetString("cid")
	utils.DeleteCollect(id)
	c.Get()
}
