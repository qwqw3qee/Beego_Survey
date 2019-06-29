package controllers

import (
	"beego_survey/models"
	"beego_survey/severio"
	"beego_survey/utils"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

const (
	headPrefix  = "head_"
	paperPrefix = "paper_"
)

//Base 基础控制器
type Base struct {
	beego.Controller
	FlashBag         *beego.FlashData
	redirectURL      string
	CurrentLoginUser *models.Users
}

//Prepare 控制器预备工作
func (Base *Base) Prepare() {

	// 初始化读取Flash
	beego.ReadFromRequest(&Base.Controller)

	// 初始化Flash
	Base.FlashBag = beego.NewFlash()
	// 自动读取当前登陆用户
	isLogin := false
	loginUserID, _ := Base.GetSession("login_user_id").(int64)
	loginUserSign, _ := Base.GetSession("login_user_sign").(string)
	if loginUserID > 0 {
		user, err := models.FindUserById(loginUserID)
		if err == nil && utils.AuthSignCheck(user.ID, user.Password, loginUserSign) == true {
			Base.CurrentLoginUser = user
			isLogin = true
		}
	}
	if !isLogin {
		loginUserID, _ = strconv.ParseInt(Base.Ctx.GetCookie("login_user_id"), 10, 64)
		loginUserSign = Base.Ctx.GetCookie("login_user_sign")
		if loginUserID > 0 {
			user, err := models.FindUserById(loginUserID)
			if err == nil && utils.AuthSignCheck(user.ID, user.Password, loginUserSign) == true {
				Base.CurrentLoginUser = user
				isLogin = true
				Base.SetSession("login_user_id", user.ID)
				Base.SetSession("login_user_sign", utils.AuthSign(user.ID, user.Password))
				Base.FlashSuccess("欢迎回来！<br>您上次登录的时间为：" + utils.TimeDiffForHumans(models.UpdateLoginTime(Base.CurrentLoginUser)))
			}
		}
	}
	//tm := lib.Time{Time: time.Now()}
	//beego.Info(tm.String())
	Base.Data["IsLogin"] = isLogin
	Base.Data["user"] = Base.CurrentLoginUser
	//beego.Info(Base.CurrentLoginUser.RegisterTime)
	Base.Data["PageTitle"] = "问卷系统"

}

//FlashSuccess 保存成功的Flash信息
func (Base *Base) FlashSuccess(message string) {
	Base.FlashBag.Notice(message)
	Base.FlashBag.Store(&Base.Controller)
}

//FlashError 保存失败的Flash信息
func (Base *Base) FlashError(message string) {
	Base.FlashBag.Error(message)
	Base.FlashBag.Store(&Base.Controller)
}

//RedirectTo 重定向
func (Base *Base) RedirectTo(url string) {
	Base.Redirect(url, 302)
	Base.StopRun()
}

// ajax错误返回
func (Base *Base) AjaxError(message string, code uint16) {
	res := severio.ErrorResponseJson{message, code}
	Base.Data["json"] = res
	Base.ServeJSON()
	Base.StopRun()
}

// ajax成功返回
func (Base *Base) AjaxSuccess(message string, data interface{}) {
	res := severio.SuccessResponseJson{message, 0, data}
	Base.Data["json"] = res
	Base.ServeJSON()
	Base.StopRun()
}

//ErrorHandler 抛出500
func (Base *Base) ErrorHandler(err error) {
	logs.Info(err)
	Base.Abort("500")
	Base.StopRun()
}

//Back 跳转到前一页
func (Base *Base) Back() {
	Base.RedirectTo(Base.Ctx.Request.Referer())
	Base.StopRun()
}
