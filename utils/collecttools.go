package utils

import (
	"archive/zip"
	"encoding/json"
	"goflume/conf"
	"goflume/models"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

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
	SaveText(filepath.Join(conf.FlumeConfPath, c.ID+".conf"), c.Setting)
	SaveAsJSON(filepath.Join(conf.CollectPath, c.ID+".json"), c)
}

//DeleteCollect 删除采集配置
func DeleteCollect(id string) {
	logs.Info("delete collect " + id)
	DeleteFile(filepath.Join(conf.CollectPath, id+".json"))
}

//PackageCollect 打包采集配置
func PackageCollect(id string) string {
	logs.Info("Create collect package for " + id)
	flumeConfPath := filepath.Join(conf.FlumeConfPath, id+".conf")
	confContent, _ := ioutil.ReadFile(flumeConfPath)
	setting := string(confContent)
	r, _ := regexp.Compile("(sources|sinks)\\.([^.]+)\\.")
	var files [][]string
	for _, s := range strings.Split(setting, "\n") {
		if !strings.HasPrefix(s, "#") && len(s) > 2 {
			nameMatch := r.FindStringSubmatch(s)
			if len(nameMatch) > 2 {
				//获取配置的source/sink名称
				name := nameMatch[2]
				//获取文件路径
				subIndex := strings.Index(s, "=")
				if -1 != subIndex {
					value := strings.Trim(s[subIndex+1:], " ")
					value = strings.ReplaceAll(value, "\r", "")
					if value[0] == '/' ||
						value[0] == 'c' ||
						value[0] == 'd' ||
						value[0] == 'e' ||
						value[0] == 'f' ||
						value[0] == 'g' ||
						value[0] == 'C' ||
						value[0] == 'D' ||
						value[0] == 'E' ||
						value[0] == 'F' ||
						value[0] == 'G' {
						if FileExist(value) {
							files = append(files, []string{name, value})
						}
					}
				}
			}
		}
	}

	zipPath := filepath.Join(os.TempDir(), "flume-"+id+"-"+FormatTime(time.Now(), "yyyyMMddHHmmss")+".zip")
	logs.Info("Create collect package at " + zipPath)
	file, _ := os.Create(zipPath)
	defer file.Close()
	writer := zip.NewWriter(file)
	defer writer.Close()

	//压缩配置文件
	confstat, _ := os.Stat(flumeConfPath)
	header, _ := zip.FileInfoHeader(confstat)
	header.Name = confstat.Name()
	src, _ := os.Open(flumeConfPath)
	dst, _ := writer.CreateHeader(header)
	io.Copy(dst, src)
	src.Close()

	//压缩配置文件中引用的文件
	for _, fileinfo := range files {
		filestat, _ := os.Stat(fileinfo[1])
		header, _ := zip.FileInfoHeader(filestat)
		header.Name = fileinfo[0] + "/" + filestat.Name()

		src, _ := os.Open(fileinfo[1])
		dst, _ := writer.CreateHeader(header)

		io.Copy(dst, src)
		src.Close()
	}

	return zipPath
}
