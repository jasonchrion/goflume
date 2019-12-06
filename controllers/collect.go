package controllers

import (
	"goflume/models"
	"goflume/utils"
	"strings"

	"github.com/astaxie/beego"
)

//CollectorController 控制器
type CollectorController struct {
	beego.Controller
}

//Get 采集器页面
func (c *CollectorController) Get() {
	c.Data["collects"] = utils.LoadCollector()
	c.TplName = "collect.tpl"
}

//Update 采集器修改页面
func (c *CollectorController) Update() {
	id := c.GetString("cid")
	c.Data["c"] = utils.GetCollector(id)
	c.Data["templates"] = utils.LoadTemplate()
	c.TplName = "collectForm.tpl"
}

//New 采集器新建页面
func (c *CollectorController) New() {
	c.Data["templates"] = utils.LoadTemplate()
	c.TplName = "collectForm.tpl"
}

//Save 保存采集器
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

//Delete 删除采集器
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

//State 采集器状态
func (c *CollectorController) State() {
	var runStates = utils.GetRunStateMap()
	var states []models.CollectorStateInfo
	for _, collector := range utils.LoadCollector() {
		state := runStates[collector.ID]
		cmd := utils.GetStartCommand(collector)
		states = append(states, models.CollectorStateInfo{ID: collector.ID,
			Name:       collector.Name,
			CMD:        cmd.Bin + " " + strings.Join(cmd.Args, " "),
			State:      state,
			CreateTime: collector.CreateTime})
	}
	utils.SortCollectorState(states)
	c.Data["states"] = states
	c.TplName = "state.tpl"
}

//Start 启动采集器
func (c *CollectorController) Start() {
	id := c.GetString("cid")
	utils.StartCollector(id)
	c.State()
}

//Stop 关闭采集器
func (c *CollectorController) Stop() {
	id := c.GetString("cid")
	utils.StopCollector(id)
	c.State()
}
