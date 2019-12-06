package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"goflume/models"
	"io/ioutil"
	"os"
	"runtime"
	"sort"
	"strings"
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
	return FormatTime(time.Now())
}

//FormatTime 时间格式化yyyy-MM-dd HH:mm:ss
func FormatTime(t time.Time) string {
	return FormatTimeByLayout(t, "yyyy-MM-dd HH:mm:ss")
}

//FormatTimeByLayout 时间格式化
func FormatTimeByLayout(t time.Time, format string) string {
	f := format
	f = strings.Replace(f, "yyyy", "2006", 1)
	f = strings.Replace(f, "MM", "01", 1)
	f = strings.Replace(f, "dd", "02", 1)
	f = strings.Replace(f, "HH", "15", 1)
	f = strings.Replace(f, "mm", "04", 1)
	f = strings.Replace(f, "ss", "05", 1)
	return t.Format(f)
}

//ParseTime 转换时间2006-01-02 15:04:05
func ParseTime(t string) time.Time {
	return ParseTimeByLayout("2006-01-02 15:04:05", t)
}

//ParseTimeByLayout 转换时间layout
func ParseTimeByLayout(layout string, t string) time.Time {
	time, err := time.Parse(layout, t)
	if nil != err {
		logs.Error(err)
	}
	return time
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

//SaveText 保存文本
func SaveText(path string, content string) {
	ioutil.WriteFile(path, []byte(content), 0666)
}

//Md5 字符串转md5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//DeleteFile 删除文件
func DeleteFile(path string) {
	if FileExist(path) {
		os.Remove(path)
	}
}

//IsOnWindows 是否在windows上运行
func IsOnWindows() bool {
	return runtime.GOOS == "windows"
}

//SortByTemplateCreateTime 模板根据时间排序
type SortByTemplateCreateTime []models.TemplateInfo

func (a SortByTemplateCreateTime) Len() int      { return len(a) }
func (a SortByTemplateCreateTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByTemplateCreateTime) Less(i, j int) bool {
	return ParseTime(a[i].CreateTime).Unix() > ParseTime(a[j].CreateTime).Unix()
}

//SortTemplate 模板排序
func SortTemplate(tis SortByTemplateCreateTime) {
	sort.Stable(tis)
}

//SortByCollectCreateTime 采集器根据时间排序
type SortByCollectCreateTime []models.CollectInfo

func (a SortByCollectCreateTime) Len() int      { return len(a) }
func (a SortByCollectCreateTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByCollectCreateTime) Less(i, j int) bool {
	return ParseTime(a[i].CreateTime).Unix() > ParseTime(a[j].CreateTime).Unix()
}

//SortCollector 采集器排序
func SortCollector(cis SortByCollectCreateTime) {
	sort.Stable(cis)
}

//SortByFileUpdateTime 文件根据时间排序
type SortByFileUpdateTime []models.FileInfo

func (a SortByFileUpdateTime) Len() int      { return len(a) }
func (a SortByFileUpdateTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByFileUpdateTime) Less(i, j int) bool {
	return a[i].UpdateTimeInt > a[j].UpdateTimeInt
}

//SortFile 文件排序
func SortFile(cis SortByFileUpdateTime) {
	sort.Stable(cis)
}

//SortByCollectorState 采集器根据时间和状态排序
type SortByCollectorState []models.CollectorStateInfo

func (a SortByCollectorState) Len() int      { return len(a) }
func (a SortByCollectorState) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByCollectorState) Less(i, j int) bool {
	if a[i].State.Run == 1 && a[j].State.Run == 1 {
		return ParseTime(a[i].CreateTime).Unix() > ParseTime(a[j].CreateTime).Unix()
	} else if a[i].State.Run == 2 && a[j].State.Run == 2 {
		return ParseTime(a[i].CreateTime).Unix() > ParseTime(a[j].CreateTime).Unix()
	} else if a[i].State.Run == 1 {
		return true
	} else if a[j].State.Run == 1 {
		return false
	} else if a[i].State.Run == 2 {
		return true
	} else if a[j].State.Run == 2 {
		return false
	}
	return ParseTime(a[i].CreateTime).Unix() > ParseTime(a[j].CreateTime).Unix()
}

//SortCollectorState 采集器运行状态排序
func SortCollectorState(cis SortByCollectorState) {
	sort.Stable(cis)
}
