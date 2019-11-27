package routers

import (
	"goflume/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/index", &controllers.MainController{}, "*:Index")

	pages := []string{"/info", "/state", "/tourist", "/collect", "/file", "/flume", "/filewatcher"}

	for i := range pages {
		beego.Router(pages[i], &controllers.MainController{})
	}

	beego.Router("/template", &controllers.TemplateController{})
	beego.Router("/template/update", &controllers.TemplateController{}, "*:Update")
	beego.Router("/template/new", &controllers.TemplateController{}, "*:New")
	beego.Router("/template/save", &controllers.TemplateController{}, "*:Save")
	beego.Router("/template/delete", &controllers.TemplateController{}, "*:Delete")

}
