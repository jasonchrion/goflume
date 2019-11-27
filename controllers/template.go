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

// Get 响应请求
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
	var t models.TemplateInfo
	c.ParseForm(&t)
	if "" == t.ID {
		t.CreateTime = utils.GetTimeNow()
		t.UpdateTime = t.CreateTime
		t.ID = utils.Md5(t.Name + t.DESC + t.CreateTime)
	} else {
		t.UpdateTime = utils.GetTimeNow()
	}
	utils.SaveTemplate(t)
	c.Get()
}

//Delete 响应请求
func (c *TemplateController) Delete() {
	id := c.GetString("tid")
	utils.DeleteTemplate(id)
	c.Get()
}
