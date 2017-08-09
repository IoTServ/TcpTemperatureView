package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}



type NowController struct {
	beego.Controller
}

func (c *NowController) Get() {
	c.TplName = "now.html"
}