package main

import (
	_ "tempview/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_"tempview/db"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/temp")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
