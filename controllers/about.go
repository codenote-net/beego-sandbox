package controllers

import (
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Prepare() {
	c.EnableXSRF = false
}

func (c *AboutController) Get() {
	c.Data["Website"] = "About page"
	c.Data["Email"] = "about@example.com"
	c.TplName = "index.tpl"
}
