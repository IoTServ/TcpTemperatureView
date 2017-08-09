package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["tempview/controllers:DEVICEController"] = append(beego.GlobalControllerRouter["tempview/controllers:DEVICEController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["tempview/controllers:DEVICEController"] = append(beego.GlobalControllerRouter["tempview/controllers:DEVICEController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
