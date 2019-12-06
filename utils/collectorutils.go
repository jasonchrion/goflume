package utils

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"goflume/conf"
	"goflume/models"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/astaxie/beego/logs"
)

var (
	//保存采集器端口
	portMap = map[string]int{}
	//保存采集器启动命令
	commandMap = map[string]StartCommand{}
)

//StartCommand 启动命令
type StartCommand struct {
	Bin  string
	Args []string
}

//LoadCollector 加载采集器
func LoadCollector() []models.CollectInfo {
	var cs = []models.CollectInfo{}

	fis, err := ioutil.ReadDir(conf.CollectorPath)

	if nil != err {
		logs.Error(err)
		return cs
	}

	for _, fi := range fis {
		t := GetCollectorByName(fi.Name())
		cs = append(cs, t)
	}

	SortCollector(cs)

	return cs
}

//GetCollectorByName 获取模板信息
func GetCollectorByName(name string) models.CollectInfo {
	body, err := ioutil.ReadFile(filepath.Join(conf.CollectorPath, name))
	var c models.CollectInfo
	if nil != err {
		return c
	}
	json.Unmarshal(body, &c)
	return c
}

//GetCollector 获取采集配置
func GetCollector(id string) models.CollectInfo {
	return GetCollectorByName(id + ".json")
}

//SaveCollector 保存采集配置
func SaveCollector(c models.CollectInfo) {
	logs.Info("save collector " + c.ID)
	SaveText(filepath.Join(conf.FlumeConfPath, c.ID+".conf"), c.Setting)
	SaveAsJSON(filepath.Join(conf.CollectorPath, c.ID+".json"), c)
}

//DeleteCollector 删除采集配置
func DeleteCollector(id string) {
	logs.Info("delete collector " + id)
	DeleteFile(filepath.Join(conf.CollectorPath, id+".json"))
}

//PackageCollector 打包采集配置
func PackageCollector(id string) string {
	logs.Info("create collector package for " + id)
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

	zipPath := filepath.Join(os.TempDir(), "flume-"+id+"-"+FormatTimeByLayout(time.Now(), "yyyyMMddHHmmss")+".zip")
	logs.Info("create collect package at " + zipPath)
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

//GetStartCommandByID 根据id获取启动命令
func GetStartCommandByID(id string) StartCommand {
	collector := GetCollector(id)
	return GetStartCommand(collector)
}

//GetStartCommand 获取启动命令
func GetStartCommand(collector models.CollectInfo) StartCommand {
	startCmd := commandMap[collector.ID]

	if "" != startCmd.Bin {
		return startCmd
	}

	r, _ := regexp.Compile("n?([a-zA-Z0-9]+).sources")
	match := r.FindStringSubmatch(collector.Setting)
	name := string(match[1])

	configFile := filepath.Join(conf.FlumeConfPath, collector.ID+".conf")

	port := strconv.Itoa(GetMetricPort(collector.ID))

	var start string
	var args []string
	if IsOnWindows() {
		start = filepath.Join(conf.FlumeBinPath, "flume-ng.cmd")
		jvmArgs := []string{"flumeCid=" + collector.ID + " -Xmx" + collector.MemSize + "m -Xms" + collector.MemSize + "m",
			";flume.monitoring.type=http",
			";flume.monitoring.port=" + port}
		args = []string{"agent",
			"--name",
			name,
			"--conf",
			conf.FlumeConfPath,
			"--conf-file",
			configFile,
			"--property",
			strings.Join(jvmArgs, "")}
	} else {
		start = filepath.Join(conf.FlumeBinPath, "flume-ng")
		args = []string{"agent",
			"--name",
			name,
			"--conf",
			conf.FlumeConfPath,
			"--conf-file",
			configFile,
			"--no-reload-conf",
			"-DflumeCid=" + collector.ID,
			"-Xmx" + collector.MemSize + "m",
			"-Xms" + collector.MemSize + "m",
			"-Dflume.monitoring.type=http",
			"-Dflume.monitoring.port=" + port}
	}
	commandMap[collector.ID] = StartCommand{Bin: start, Args: args}
	return commandMap[collector.ID]
}

//GetListenMetricPort 获取运行中的监控端口
func GetListenMetricPort(id string) int {
	return portMap[id]
}

//GetMetricPort 获取监控端口
func GetMetricPort(id string) int {
	port := portMap[id]
	if 0 != port {
		return port
	}
	getPort := true
	for getPort {
		port = 40000 + rand.Intn(10000)
		getPort = false
		for _, p := range portMap {
			if p == port {
				getPort = true
				break
			}
		}
	}
	portMap[id] = port
	return port
}

//GetRunStateMap 获取采集器运行状态
//
//返回cid:[cid,port,pid]
func GetRunStateMap() map[string]models.CollectorRunInfo {
	if IsOnWindows() {
		return GetRunStateMapOnWindows()
	} else {
		return GetRunStateMapOnLinux()
	}
}

//GetRunStateMapOnLinux linux下获取运行状态
func GetRunStateMapOnLinux() map[string]models.CollectorRunInfo {
	var stateMap = map[string]models.CollectorRunInfo{}
	return stateMap
}

//GetRunStateMapOnWindows windows下获取运行状态
func GetRunStateMapOnWindows() map[string]models.CollectorRunInfo {
	var stateMap = map[string]models.CollectorRunInfo{}
	cmd := exec.Command("cmd", "/C", "wmic process where caption='java.exe' get commandline,processid /value")
	stdout, err := cmd.StdoutPipe()
	defer stdout.Close()
	if nil != err {
		logs.Error(err)
		return stateMap
	}
	err2 := cmd.Start()
	if nil != err2 {
		logs.Error(err2)
		return stateMap
	}
	reader := bufio.NewReader(stdout)

	var states [][]string
	cidReg, _ := regexp.Compile("-DflumeCid=([a-zA-Z0-9]{32})")
	portReg, _ := regexp.Compile("-Dflume.monitoring.port=([0-9]{5})")
	pidReg, _ := regexp.Compile("[0-9]+")
	skip := false
	for {
		line, err3 := reader.ReadString('\n')
		if err3 != nil {
			break
		}
		if strings.HasPrefix(line, "CommandLine=") {
			cidFind := cidReg.FindStringSubmatch(line)
			if 0 == len(cidFind) {
				skip = true
				continue
			}
			port := portMap[cidFind[1]]
			if 0 != port {
				states = append(states, []string{cidFind[1], strconv.Itoa(port)})
			} else {
				portFind := portReg.FindStringSubmatch(line)
				if 0 == len(portFind) {
					skip = true
					continue
				}
				states = append(states, []string{cidFind[1], portFind[1]})
			}
		} else if strings.HasPrefix(line, "ProcessId=") {
			if skip {
				skip = false
				continue
			} else {
				pidFind := pidReg.FindString(line)
				states[len(states)-1] = append(states[len(states)-1], pidFind)
			}
		}
	}
	for _, state := range states {
		a1, _ := strconv.Atoi(state[1])
		a2, _ := strconv.Atoi(state[2])
		a3 := models.CollectorRunInfo{ID: state[0],
			Port: a1,
			PID:  a2,
			Run:  1}
		stateMap[state[0]] = a3
		//更新采集器监控端口
		portMap[a3.ID] = a3.Port
	}
	return stateMap
}

//GetRunInfo 获取采集器运行信息
func GetRunInfo(id string) models.CollectorRunInfo {
	return GetRunStateMap()[id]
}

//StartCollector 启动采集器
func StartCollector(id string) {
	if checkStart(id) {
		return
	}
	startCmd := GetStartCommandByID(id)
	if "" != startCmd.Bin {
		logs.Info("start collector " + id)
		if IsOnWindows() {
			go func() {
				//windows下启动
				cmd := exec.Command(startCmd.Bin, startCmd.Args...)
				stdout, err := cmd.StderrPipe()
				defer stdout.Close()
				if nil != err {
					logs.Error(err)
					return
				}
				err2 := cmd.Start()
				if nil != err2 {
					logs.Error(err2)
					return
				}
				reader := bufio.NewReader(stdout)
				for {
					_, err3 := reader.ReadString('\n')
					if err3 != nil {
						break
					}
				}
				logs.Info("close collector error stream " + id)
			}()
			waitFor(60, true, id)
		} else {
			//linux下启动
			err := syscall.Exec(startCmd.Bin, startCmd.Args, os.Environ())
			if nil != err {
				logs.Error(err)
			}
		}
	}
}

//StopCollector 关闭采集器
func StopCollector(id string) {
	runInfo := GetRunInfo(id)
	logs.Info("close collector", runInfo)
	if 0 != runInfo.PID {
		if IsOnWindows() {
			//windows下关闭
			cmd := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(runInfo.PID))
			cmd.Start()
		} else {
			//linux下关闭

		}
		//清除监控端口
		portMap[id] = 0
		waitFor(60, false, id)
	}
}

//检查是否启动
func checkStart(id string) bool {
	return GetRunInfo(id).PID != 0
}

//等待启动/关闭
func waitFor(waitSecond int, isStartAction bool, id string) {
	time.Sleep(2 * time.Second)
	for wait := waitSecond - 2; wait > 0; wait-- {
		if checkStart(id) == isStartAction {
			break
		}
		time.Sleep(5 * time.Second)
	}
}
