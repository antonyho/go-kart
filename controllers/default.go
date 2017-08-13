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
// @Success 200 string status
// @Failure 400 Bad request
// @Accept json
// @router / [get]
func (c *HealthController) Status() {
	status := "Running"
	c.Data["json"] = &status
	c.ServeJSON()
}
