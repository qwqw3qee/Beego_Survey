package controllers

import "github.com/astaxie/beego"

//IndexController 主页
type IndexController struct {
	Base
}

// @router / [get]
func (c *IndexController) Index() {
	c.redirectURL = beego.URLFor("IndexController.Index")
	c.TplName = "index.html"
	c.Data["PageTitle"] = "问卷系统"

}
