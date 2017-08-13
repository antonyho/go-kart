package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/antonyho/go-kart/controllers:HealthController"] = append(beego.GlobalControllerRouter["github.com/antonyho/go-kart/controllers:HealthController"],
		beego.ControllerComments{
			Method: "Status",
			Router: `/status`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
