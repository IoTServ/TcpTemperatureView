package routers

import (
	"tempview/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/now", &controllers.NowController{})
	ns := beego.NewNamespace("/api",

		beego.NSNamespace("/data",
			beego.NSInclude(
				&controllers.DEVICEController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
