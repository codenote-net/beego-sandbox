package routers

import (
	"github.com/astaxie/beego"
	"github.com/codenote-net/beego-sandbox/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{})
}
