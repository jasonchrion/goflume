package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"goflume/models"
	"io/ioutil"
	"os"
	"sort"
	"time"

	"github.com/astaxie/beego/logs"
)

//FileExist 判断文件是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if nil == err {
		return true
	}
	return false
}

//CreateDir 创建目录
func CreateDir(paths ...string) {
	for _, path := range paths {
		if FileExist(path) {
			continue
		}
		err := os.MkdirAll(path, 0666)
		if nil != err {
			logs.Error(err)
		}
	}
}

//GetTimeNow 返回格式化当前时间
func GetTimeNow() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}

//SaveAsJSON 保存为json文本
func SaveAsJSON(path string, obj interface{}) {
	data, err := json.Marshal(obj)
	if nil != err {
		logs.Error(err)
		return
	}
	ioutil.WriteFile(path, data, 0666)
}

//Md5 字符串转md5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//DeleteFile 删除文件
func DeleteFile(path string) {
	os.Remove(path)
}

//SortByTemplateCreateTime 根据时间排序
type SortByTemplateCreateTime []models.TemplateInfo

func (a SortByTemplateCreateTime) Len() int           { return len(a) }
func (a SortByTemplateCreateTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByTemplateCreateTime) Less(i, j int) bool { return a[i].CreateTime > a[j].CreateTime }

//SortTemplate 模板排序
func SortTemplate(tis SortByTemplateCreateTime) {
	sort.Stable(tis)
}
