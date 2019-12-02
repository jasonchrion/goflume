package main

import (
	"goflume/conf"
	_ "goflume/routers"
	"goflume/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	start()
}

func start() {
	beego.SetLogger(logs.AdapterFile, `{"filename":"`+conf.UILogPath+`"}`)

	checkdir()

	beego.Run()
}

func checkdir() {
	utils.CreateDir(conf.CollectPath, conf.TemplatePath, conf.FileStorePath, conf.FlumePath, conf.FlumeConfPath)
	utils.CreateSubDir()
}
