package controllers

import (
	"goflume/conf"
	"goflume/utils"
	"path/filepath"

	"github.com/astaxie/beego"
)

//FileController 页面跳转
type FileController struct {
	beego.Controller
}

//Get Get页面
func (c *FileController) Get() {
	c.Data["files"] = utils.GetFileInfos()
	c.TplName = "file.tpl"
}

//Download 下载文件
func (c *FileController) Download() {
	name := c.GetString("name")
	c.Ctx.Output.Download(filepath.Join(conf.FileStorePath, name))
}

//Delete 删除文件
func (c *FileController) Delete() {
	name := c.GetString("name")
	utils.DeleteFile(filepath.Join(conf.FileStorePath, name))
	c.Get()
}

//Upload 上传文件
func (c *FileController) Upload() {
	f, h, err := c.GetFile("inputfile")
	defer f.Close()
	if nil == err {
		c.SaveToFile("inputfile", utils.GetStorePath(h.Filename))
	}
	c.Get()
}

//Update 修改文件
func (c *FileController) Update() {
	name := c.GetString("name")
	c.Data["fi"] = utils.GetFileInfo(filepath.Join(conf.FileStorePath, name), true)
	c.TplName = "fileForm.tpl"
}

//Save 保存修改内容
func (c *FileController) Save() {
	name := c.GetString("path")
	content := c.GetString("content")
	utils.SaveText(filepath.Join(conf.FileStorePath, name), content)
	c.Get()
}
