package routers

import (
	"github.com/antonyho/go-kart/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
