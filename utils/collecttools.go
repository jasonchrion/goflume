package utils

import (
	"encoding/json"
	"goflume/conf"
	"goflume/models"
	"io/ioutil"
	"path/filepath"

	"github.com/astaxie/beego/logs"
)

//LoadCollect 加载模板
func LoadCollect() []models.CollectInfo {
	var cs = []models.CollectInfo{}

	fis, err := ioutil.ReadDir(conf.CollectPath)

	if nil != err {
		logs.Error(err)
		return cs
	}

	for _, fi := range fis {
		t := GetCollectByName(fi.Name())
		cs = append(cs, t)
	}

	SortCollect(cs)

	return cs
}

//GetCollectByName 获取模板信息
func GetCollectByName(name string) models.CollectInfo {
	body, err := ioutil.ReadFile(filepath.Join(conf.CollectPath, name))
	var c models.CollectInfo
	if nil != err {
		logs.Error(err)
		return c
	}
	json.Unmarshal(body, &c)
	return c
}

//GetCollect 获取采集配置
func GetCollect(id string) models.CollectInfo {
	return GetCollectByName(id + ".json")
}

//SaveCollect 保存采集配置
func SaveCollect(c models.CollectInfo) {
	logs.Info("save collect " + c.ID)
	SaveAsJSON(filepath.Join(conf.CollectPath, c.ID+".json"), c)
}

//DeleteCollect 删除采集配置
func DeleteCollect(id string) {
	logs.Info("delete collect " + id)
	DeleteFile(filepath.Join(conf.CollectPath, id+".json"))
}
