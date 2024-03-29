package routers

import (
	"goflume/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/index", &controllers.MainController{}, "*:Index")

	pages := []string{"/info", "/flume", "/filewatcher"}

	for i := range pages {
		beego.Router(pages[i], &controllers.MainController{})
	}

	beego.Router("/template", &controllers.TemplateController{})
	beego.Router("/template/update", &controllers.TemplateController{}, "*:Update")
	beego.Router("/template/new", &controllers.TemplateController{}, "*:New")
	beego.Router("/template/save", &controllers.TemplateController{}, "*:Save")
	beego.Router("/template/delete", &controllers.TemplateController{}, "*:Delete")

	beego.Router("/tourist", &controllers.TouristController{})
	beego.Router("/tourist/page", &controllers.TouristController{}, "*:Page")

	beego.Router("/collect", &controllers.CollectorController{})
	beego.Router("/collect/update", &controllers.CollectorController{}, "*:Update")
	beego.Router("/collect/new", &controllers.CollectorController{}, "*:New")
	beego.Router("/collect/save", &controllers.CollectorController{}, "*:Save")
	beego.Router("/collect/delete", &controllers.CollectorController{}, "*:Delete")
	beego.Router("/collect/package", &controllers.CollectorController{}, "*:Package")
	beego.Router("/collect/start", &controllers.CollectorController{}, "*:Start")
	beego.Router("/collect/stop", &controllers.CollectorController{}, "*:Stop")
	beego.Router("/state", &controllers.CollectorController{}, "*:State")

	beego.Router("/file", &controllers.FileController{})
	beego.Router("/file/download", &controllers.FileController{}, "*:Download")
	beego.Router("/file/delete", &controllers.FileController{}, "*:Delete")
	beego.Router("/file/upload", &controllers.FileController{}, "*:Upload")
	beego.Router("/file/update", &controllers.FileController{}, "*:Update")
	beego.Router("/file/save", &controllers.FileController{}, "*:Save")

	beego.Router("/ws/uilog", &controllers.WebSocketController{}, "*:UILog")
	beego.Router("/ws/log", &controllers.WebSocketController{}, "*:CollectorLog")
	beego.Router("/ws/jmx/metric", &controllers.WebSocketController{}, "*:Metric")

}
