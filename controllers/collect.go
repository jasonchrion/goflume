package controllers

import (
	"goflume/models"
	"goflume/utils"

	"github.com/astaxie/beego"
)

//CollectorController 控制器
type CollectorController struct {
	beego.Controller
}

//Get 响应请求
func (c *CollectorController) Get() {
	c.Data["collects"] = utils.LoadCollector()
	c.TplName = "collect.tpl"
}

//Update 响应请求
func (c *CollectorController) Update() {
	id := c.GetString("cid")
	c.Data["c"] = utils.GetCollector(id)
	c.Data["templates"] = utils.LoadTemplate()
	c.TplName = "collectForm.tpl"
}

//New 响应请求
func (c *CollectorController) New() {
	c.Data["templates"] = utils.LoadTemplate()
	c.TplName = "collectForm.tpl"
}

//Save 响应请求
func (c *CollectorController) Save() {
	var ci models.CollectInfo
	c.ParseForm(&ci)
	if "" == ci.ID {
		ci.CreateTime = utils.GetTimeNow()
		ci.UpdateTime = ci.CreateTime
		ci.ID = utils.Md5("C-" + ci.Name + ci.DESC + ci.CreateTime)
	} else {
		ci.UpdateTime = utils.GetTimeNow()
	}
	ci2 := utils.GetCollector(ci.ID)
	if "" != ci2.ID {
		if "" == ci2.CreateTime {
			ci.CreateTime = ci.UpdateTime
		} else {
			ci.CreateTime = ci2.CreateTime
		}
	}
	utils.SaveCollector(ci)
	c.Get()
}

//Delete 响应请求
func (c *CollectorController) Delete() {
	id := c.GetString("cid")
	utils.DeleteCollector(id)
	c.Get()
}

//Package 配置打包
func (c *CollectorController) Package() {
	id := c.GetString("cid")
	zipPath := utils.PackageCollector(id)
	c.Ctx.Output.Download(zipPath)
}
