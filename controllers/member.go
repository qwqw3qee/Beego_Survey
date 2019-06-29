package controllers

import (
	"beego_survey/models"
	"beego_survey/utils"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type MemberController struct {
	Base
}

func (c *MemberController) Prepare() {
	c.Base.Prepare()
	if c.CurrentLoginUser == nil {
		c.RedirectTo(beego.URLFor("UserController.Login"))
	}

}

// @router /member/logout [get]
func (c *MemberController) Logout() {
	c.CurrentLoginUser = nil
	c.Ctx.SetCookie("login_user_id", "", -3600)
	c.Ctx.SetCookie("login_user_sign", "", -3600)
	c.DelSession("login_user_id")
	c.DelSession("login_user_sign")
	c.FlashSuccess("已安全退出")
	c.RedirectTo("/")
}

// @router /member/headimg [get]
func (c *MemberController) HeadImg() {
	url := "/headimg/" + strconv.FormatInt(c.CurrentLoginUser.ID, 10)
	beego.Info(url)
	c.RedirectTo(url)
}

// @router /member/info [get]
func (c *MemberController) MemberIndex() {
	c.TplName = "member_info.html"
	c.Data["PageTitle"] = "我的信息"
	c.Data["MemberInfo"] = true
	papers, _ := models.PapersCountByPerson(c.CurrentLoginUser.ID)
	answers, _ := models.CountRecordsByAnswerUserID(c.CurrentLoginUser.ID)
	c.Data["papers"] = papers
	c.Data["answers"] = answers
}

// @router /member/change_password [post]
func (c *MemberController) PasswordChangeHandler() {
	oldPwd := c.Input().Get("oldPwd")
	newPwd1 := c.Input().Get("newPwd1")
	newPwd2 := c.Input().Get("newPwd2")
	beego.Info(oldPwd)
	if utils.SHA256Encode(oldPwd) != c.CurrentLoginUser.Password {
		c.AjaxError("原始密码错误", 1)
	}
	if oldPwd == newPwd1 {
		c.AjaxError("新密码与原始密码相同", 2)
	}
	if newPwd1 != newPwd2 {
		c.AjaxError("两次输入不一致", 3)
	}
	if models.PasswordUpdate(c.CurrentLoginUser, newPwd1) != nil {
		c.AjaxError("更新数据库出错", 4)
	}
	c.SetSession("login_user_sign", utils.AuthSign(c.CurrentLoginUser.ID, c.CurrentLoginUser.Password))
	c.AjaxSuccess("修改成功", "engoy")

}

// @router /member/change_nickname [post]
func (c *MemberController) NickNameChangeHandler() {
	newNickName := strings.TrimSpace(c.Input().Get("newNick"))
	if len(newNickName) == 0 {
		c.AjaxError("昵称为空，请重新输入", 1)
	}
	if models.NickNameUpdate(c.CurrentLoginUser, newNickName) != nil {
		c.AjaxError("更新数据库失败", 2)
	}
	res := make(map[string]string)
	res["nick"] = newNickName
	c.AjaxSuccess("修改成功", res)
}

// @router /member/do [get]
func (c *MemberController) MemberDo() {
	c.TplName = "member_do.html"
	c.Data["PageTitle"] = "我的参与"
	c.Data["MemberDo"] = true
}

// @router /member/publish [get]
func (c *MemberController) MemberPublish() {
	c.TplName = "member_publish.html"
	c.Data["PageTitle"] = "我的发布"
	c.Data["MemberPublish"] = true
}

// @router /member/publish/dolist [get]
func (c *MemberController) PublishDolist() {
	paperID, err := c.GetInt64("paperid", 0)
	if err != nil {
		c.FlashError("参数不合法")
	}
	c.TplName = "member_publish_dolist.html"
	c.Data["PageTitle"] = "记录详情"
	c.Data["MemberPublish"] = true
	c.Data["queryPaperID"] = paperID
}
