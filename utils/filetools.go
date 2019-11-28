package utils

import (
	"goflume/conf"
	"goflume/models"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	//SubDirCount 下级文件夹个数
	SubDirCount = 32
	//SubDirNames 下级文件夹名称
	SubDirNames []string
)

//GetFileInfos 获取上传文件信息
func GetFileInfos() []models.FileInfo {
	var files []models.FileInfo
	dirs := ReadDirs(conf.FileStorePath)
	for _, fpath := range dirs {
		files = append(files, GetFileInfo(fpath, false))
	}
	SortFile(files)
	return files
}

//GetFileInfo 根据路径获取文件信息
func GetFileInfo(path string, content bool) models.FileInfo {
	fs, _ := os.Stat(path)
	path = strings.ReplaceAll(path, "\\", "/")
	fi := models.FileInfo{FileName: fs.Name(),
		FilePath:      path,
		ShortPath:     path[len(conf.FileStorePath)+1:],
		UpdateTime:    FormatTime(fs.ModTime()),
		UpdateTimeInt: fs.ModTime().Unix()}
	if content {
		c, err := ioutil.ReadFile(path)
		if nil == err {
			fi.Content = string(c)
		}
	}
	return fi
}

//CreateSubDir 创建下级目录
func CreateSubDir() {
	for i := 0; i < SubDirCount; i++ {
		hex := strconv.FormatInt(int64(i), 16)
		if 1 == len(hex) {
			hex = "0" + hex
		}
		SubDirNames = append(SubDirNames, hex)
		CreateDir(filepath.Join(conf.FileStorePath, hex))
	}
}

//ReadDirs 遍历指定目录获取全部文件路径
func ReadDirs(dir string) []string {
	var dirpaths []string
	fis, err := ioutil.ReadDir(dir)
	if nil != err {
		return dirpaths
	}
	for _, fi := range fis {
		if fi.IsDir() {
			subdirs := ReadDirs(filepath.Join(dir, fi.Name()))
			dirpaths = append(dirpaths, subdirs...)
		} else {
			dirpaths = append(dirpaths, filepath.Join(dir, fi.Name()))
		}
	}
	return dirpaths
}

//GetStorePath 获取文件存储路径
func GetStorePath(name string) string {
	var path string
	for _, n := range SubDirNames {
		path = filepath.Join(conf.FileStorePath, n, name)
		if !FileExist(path) {
			break
		}
	}
	return path
}
