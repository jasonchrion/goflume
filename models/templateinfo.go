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
