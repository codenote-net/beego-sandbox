package routers

import (
	"github.com/codenote-net/beego-sandbox/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
