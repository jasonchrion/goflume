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
