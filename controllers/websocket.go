package controllers

import (
	"bufio"
	"goflume/conf"
	"goflume/utils"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	tailWait   = 5 * time.Second
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

//WebSocketController 控制器
type WebSocketController struct {
	beego.Controller
}

//UILog 页面日志跟踪
func (c *WebSocketController) UILog() {
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, upgrader.ReadBufferSize, upgrader.WriteBufferSize)
	if nil != err {
		logs.Error(err)
		return
	}
	ws.WriteMessage(websocket.TextMessage, []byte("200"))
	go tail(ws, conf.UILogPath)
	readClient(ws, func(addr string, b []byte) {
		logs.Info(addr + ": " + string(b))
	})
}

//CollectorLog 采集器日志跟踪
func (c *WebSocketController) CollectorLog() {
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, upgrader.ReadBufferSize, upgrader.WriteBufferSize)
	if nil != err {
		logs.Error(err)
		return
	}
	ws.WriteMessage(websocket.TextMessage, []byte("200"))
	cid := c.GetString("cid")
	go tail(ws, conf.LogPath+"/"+cid+"/flume-"+cid+".log")
	readClient(ws, func(addr string, b []byte) {
		logs.Info(addr + ": " + string(b))
	})
}

//Metric 采集器数据监控
func (c *WebSocketController) Metric() {
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, upgrader.ReadBufferSize, upgrader.WriteBufferSize)
	if nil != err {
		logs.Error(err)
		return
	}
	ws.WriteMessage(websocket.TextMessage, []byte("200"))
	go readClient(ws, func(addr string, b []byte) {
		content := string(b)
		if len(content) == 32 {
			port := utils.GetListenMetricPort(content)
			if 0 != port {
				resp, err := http.Get("http://localhost:" + strconv.Itoa(port) + "/metrics")
				if nil != err {
					ws.WriteMessage(websocket.TextMessage, []byte(err.Error()))
				} else {
					body, _ := ioutil.ReadAll(resp.Body)
					msg := append([]byte(content+"#"), body...)
					ws.WriteMessage(websocket.TextMessage, msg)
				}
			}
		} else {
			logs.Info(addr, content)
		}
	})
}

func tail(ws *websocket.Conn, path string) {
	f, err := os.Open(path)
	defer f.Close()
	if nil != err {
		errMsg := path + " does not exist !!!"
		logs.Error(errMsg)
		ws.WriteMessage(websocket.TextMessage, []byte(errMsg))
		ws.Close()
		return
	}

	logs.Info("watch file " + path)
	f.Seek(-1024*2, 2)

	br := bufio.NewReader(f)

	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			time.Sleep(tailWait)
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				logs.Error("ping message error:", err)
				return
			}
			continue
		} else if err != nil {
			logs.Error("read line error:", err)
			return
		}
		ws.SetWriteDeadline(time.Now().Add(writeWait))
		if err := ws.WriteMessage(websocket.TextMessage, line); err != nil {
			logs.Error("write message error:", err)
			return
		}
	}
}

func readClient(ws *websocket.Conn, doJob func(string, []byte)) {
	address := ws.RemoteAddr().String()
	logs.Info(address + " connected")
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			logs.Error(err)
			break
		}
		doJob(address, p)
	}
}
