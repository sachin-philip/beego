package routers

import (
	"core/controllers"
	"github.com/astaxie/beego"
	// "http.ServeFile"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.SetStaticPath("/static","static")
    beego.SetStaticPath("/images","img")
	beego.SetStaticPath("/css","css")
	beego.SetStaticPath("/js","js")
    beego.Router("/next", &controllers.MainController{}, "get:Next")
}
