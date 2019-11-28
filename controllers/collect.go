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
	ci2 := utils.GetCollect(ci.ID)
	if "" != ci2.ID {
		if "" == ci2.CreateTime {
			ci.CreateTime = ci.UpdateTime
		} else {
			ci.CreateTime = ci2.CreateTime
		}
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

//Package 配置打包
func (c *CollectController) Package() {
	id := c.GetString("cid")
	zipPath := utils.PackageCollect(id)
	c.Ctx.Output.Download(zipPath)
}
