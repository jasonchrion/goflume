package models

//TemplateInfo 模板信息
type TemplateInfo struct {
	ID         string `json:"id" form:"tid"`
	Name       string `json:"name" form:"name"`
	DESC       string `json:"desc" form:"desc"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	Setting    string `json:"setting" form:"setting"`
}

//CollectInfo 模板信息
type CollectInfo struct {
	ID             string `json:"id" form:"cid"`
	Name           string `json:"name" form:"name"`
	DESC           string `json:"desc" form:"desc"`
	CreateTime     string `json:"createTime"`
	UpdateTime     string `json:"updateTime"`
	Setting        string `json:"setting" form:"setting"`
	MemSize        string `json:"memSize" form:"memSize"`
	Company        string `json:"company" form:"company"`
	Product        string `json:"product" form:"product"`
	ProductVersion string `json:"productVersion" form:"productVersion"`
}

//FileInfo 上传的配置文件信息
type FileInfo struct {
	FileName      string
	FilePath      string
	ShortPath     string
	UpdateTime    string
	UpdateTimeInt int64
	Content       string
}

//CollectorStateInfo 采集器运行状态信息
type CollectorStateInfo struct {
	ID         string           `json:"id"`
	Name       string           `json:"name"`
	State      CollectorRunInfo `json:"state"`
	CMD        string           `json:"cmd"`
	CreateTime string           `json:"createTime"`
}

//CollectorRunInfo 运行情况
type CollectorRunInfo struct {
	ID   string `json:"id"`
	Port int    `json:"port"`
	PID  int    `json:"pid"`
	//0-关闭 1-运行 2-重启
	Run int `json:"run"`
}
