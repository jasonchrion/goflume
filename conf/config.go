package conf

import "path/filepath"

var (
	//RootPath 根目录
	RootPath = "D:\\flume_ui2"
	//CollectPath collect保存目录
	CollectPath = filepath.Join(RootPath, "collect")
	//TemplatePath template保存目录
	TemplatePath = filepath.Join(RootPath, "template")
	//FilePath file保存目录
	FilePath = filepath.Join(RootPath, "file")
)
