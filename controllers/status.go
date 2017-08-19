package controllers

import (
	"github.com/astaxie/beego"
)

// Entry Point of API
type HealthController struct {
	beego.Controller
}

// @Title Service status
// @Summary Service status of this shopping cart
// @Description Query the status of this shopping cart
// @Success 200 {string} status
// @router /status [get]
func (c *HealthController) Status() {
	status := "Running"
	c.Data["json"] = map[string]string{"Status": status}
	c.ServeJSON()
}
