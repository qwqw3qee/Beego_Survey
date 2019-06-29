package controllers

import (
	"beego_survey/models"
)

type ApiMemberController struct {
	Base
}

func (c *ApiMemberController) Prepare() {
	c.Base.Prepare()
	if c.CurrentLoginUser == nil {
		c.AjaxError("请登录", 999)
	}

}

//@router /api/member/do [get]
func (c *ApiMemberController) MemberDo() {

	record, err := models.FindRecordsByAnswerUserID(c.CurrentLoginUser.ID)

	if err != nil {
		c.AjaxError("查询失败", 1)
	}
	data := map[string]interface{}{"total": len(record)}
	data["rows"] = record
	c.Data["json"] = data
	c.ServeJSON()
}

//@router /api/member/publish [get]
func (c *ApiMemberController) MemberPublish() {

	papers, err := models.FindPaperByAuthorID(c.CurrentLoginUser.ID)

	if err != nil {
		c.AjaxError("查询失败", 1)
	}
	data := map[string]interface{}{"total": len(papers)}
	data["rows"] = papers
	c.Data["json"] = data
	c.ServeJSON()
}

//@router /api/member/publish/dolist [get]
func (c *ApiMemberController) PublishDolist() {

	paperID, err := c.GetInt64("paperid", 0)
	if err != nil {
		c.AjaxError("查询参数有误", 1)
	}
	var paper *models.Papers
	paper, err = models.FindPaperByPaperID(paperID)
	if err != nil {
		c.AjaxError("查无此问卷", 2)
	}
	if c.CurrentLoginUser.UserType < 2 && paper.AuthorID.ID != c.CurrentLoginUser.ID {
		c.AjaxError("您没有权限查看记录", 3)
	}
	var record []*models.Records

	record, err = models.FindRecordsByPaperID(paperID)
	if err != nil {
		c.AjaxError("查询失败", 1)
	}

	data := map[string]interface{}{"total": len(record)}
	data["rows"] = record
	c.Data["json"] = data
	c.ServeJSON()
}
