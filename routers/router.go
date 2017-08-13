// @APIVersion 1.0.0
// @Title go-kart API
// @Description RESTful API for "go-kart" shopping cart written in Go
// @Contact ntonyworkshop@gmail.com
package routers

import (
	"github.com/antonyho/go-kart/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/status", beego.NSInclude(
			&controllers.HealthController{}),
		),
	)
    beego.AddNamespace(ns)
}
