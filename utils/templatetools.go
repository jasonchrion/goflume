package utils

import (
	"encoding/json"
	"goflume/conf"
	"goflume/models"
	"io/ioutil"
	"path/filepath"

	"github.com/astaxie/beego/logs"
)

//LoadTemplate 加载模板
func LoadTemplate() []models.TemplateInfo {
	var ts = []models.TemplateInfo{}

	fis, err := ioutil.ReadDir(conf.TemplatePath)

	if nil != err {
		logs.Error(err)
		return ts
	}

	for _, fi := range fis {
		t := GetTemplateByName(fi.Name())
		ts = append(ts, t)
	}

	SortTemplate(ts)

	return ts
}

//GetTemplateByName 获取模板信息
func GetTemplateByName(name string) models.TemplateInfo {
	body, err := ioutil.ReadFile(filepath.Join(conf.TemplatePath, name))
	var t models.TemplateInfo
	if nil != err {
		return t
	}
	json.Unmarshal(body, &t)
	return t
}

//GetTemplate 获取模板信息
func GetTemplate(id string) models.TemplateInfo {
	return GetTemplateByName(id + ".json")
}

//SaveTemplate 保存模板
func SaveTemplate(t models.TemplateInfo) {
	logs.Info("save template " + t.ID)
	SaveAsJSON(filepath.Join(conf.TemplatePath, t.ID+".json"), t)
}

//DeleteTemplate 删除模板
func DeleteTemplate(id string) {
	logs.Info("delete template " + id)
	DeleteFile(filepath.Join(conf.TemplatePath, id+".json"))
}
