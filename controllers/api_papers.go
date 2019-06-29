package controllers

import (
	"beego_survey/models"

	"github.com/astaxie/beego"
)

type ApiPapersController struct {
	Base
}

//@router /api/paper/list [get]
func (c *ApiPapersController) PaperList() {
	rows, _ := c.GetInt64("limit", 1000)
	page, _ := c.GetInt64("offset", 0)
	sortCol := c.GetString("sort", "")
	isDesc := c.GetString("sortOrder", "") == "desc"
	beego.Info(rows, page, sortCol, isDesc)
	papers := []models.Papers{}
	cols, err := models.PapersPart(&papers, rows, page, sortCol, isDesc)
	if err != nil {
		c.AjaxError("查询失败", 1)
	}
	data := map[string]interface{}{"total": cols}
	data["rows"] = papers
	c.Data["json"] = data
	c.ServeJSON()
}

