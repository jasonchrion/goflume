package conf

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/astaxie/beego"
)

var (
	//RootPath 根目录
	RootPath = "D:\\flume_ui2"
	//CollectPath collect保存目录
	CollectPath = filepath.Join(RootPath, "collect")
	//TemplatePath template保存目录
	TemplatePath = filepath.Join(RootPath, "template")
	//FileStorePath file保存目录
	FileStorePath = filepath.Join(RootPath, "file")
	//TouristConfig Tourist配置
	TouristConfig = getTouristConfig()
	//FlumePath flume根目录
	FlumePath = filepath.Join(RootPath, "flume")
	//FlumeConfPath flume配置文件目录
	FlumeConfPath = filepath.Join(FlumePath, "conf")
	//UILogPath 页面日志路径
	UILogPath = "/var/log/goflume/goflume.log"
)

//Tourist tourist配置信息
type Tourist struct {
	SourceMap  map[string]string `yaml:"source-map"`
	ChannelMap map[string]string `yaml:"channel-map"`
	SinkMap    map[string]string `yaml:"sink-map"`
}

func getTouristConfig() Tourist {
	var t Tourist
	path := filepath.Join(beego.AppPath, "/conf/tourist.yml")
	_, err := os.Stat(path)
	if nil != err {
		return t
	}
	data, err := ioutil.ReadFile(path)
	if nil != err {
		return t
	}
	yaml.Unmarshal(data, &t)
	return t
}
