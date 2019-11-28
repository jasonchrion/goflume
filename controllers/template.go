package controllers

import (
	"goflume/models"
	"goflume/utils"

	"github.com/astaxie/beego"
)

//TemplateController 控制器
type TemplateController struct {
	beego.Controller
}

//Get 响应请求
func (c *TemplateController) Get() {
	c.Data["templates"] = utils.LoadTemplate()
	c.TplName = "template.tpl"
}

//Update 响应请求
func (c *TemplateController) Update() {
	id := c.GetString("tid")
	c.Data["t"] = utils.GetTemplate(id)
	c.TplName = "templateForm.tpl"
}

//New 响应请求
func (c *TemplateController) New() {
	c.TplName = "templateForm.tpl"
}

//Save 响应请求
func (c *TemplateController) Save() {
	var ti models.TemplateInfo
	c.ParseForm(&ti)
	if "" == ti.ID {
		ti.CreateTime = utils.GetTimeNow()
		ti.UpdateTime = ti.CreateTime
		ti.ID = utils.Md5("T-" + ti.Name + ti.DESC + ti.CreateTime)
	} else {
		ti.UpdateTime = utils.GetTimeNow()
	}
	ti2 := utils.GetTemplate(ti.ID)
	if "" != ti2.ID {
		if "" == ti2.CreateTime {
			ti.CreateTime = ti.UpdateTime
		} else {
			ti.CreateTime = ti2.CreateTime
		}
	}
	utils.SaveTemplate(ti)
	c.Get()
}

//Delete 响应请求
func (c *TemplateController) Delete() {
	id := c.GetString("tid")
	utils.DeleteTemplate(id)
	c.Get()
}
