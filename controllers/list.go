package controllers

import (
	"github.com/astaxie/beego"
)

// Listing item or group API
type ListController struct {
	beego.Controller
}

// @Title List tree of group and item
// @Summary Listing with a tree of group and item
// @Description Listing with a tree of all groups and items
// @Success 200 {object} models.Group
// @router /all [get]
func (c *ListController) All() {
}
