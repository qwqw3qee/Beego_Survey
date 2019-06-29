package controllers

import (
	"beego_survey/models"
	"beego_survey/utils"
	"path"
	"strconv"
	"strings"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
)

//UserController 用户相关控制器
type UserController struct {
	Base
}

// @router /login [get]
func (c *UserController) Login() {
	if c.CurrentLoginUser != nil {
		c.RedirectTo("/")
	}
	c.TplName = "login.html"
	c.Data["PageTitle"] = "登录"
	//this.Data["PageKeywords"] = this.SettingData["SEO_INDEX_KEYWORDS"]
	//this.Data["PageDescription"] = this.SettingData["SEO_INDEX_DESCRIPTION"]
}

// @router /login [post]
func (c *UserController) LoginHandler() {
	c.redirectURL = beego.URLFor("UserController.Login")
	//loginData := fronted.UserLoginValidation{}
	//this.ValidatorAuto(&loginData)
	uname := strings.TrimSpace(c.Input().Get("uname"))
	pwd := c.Input().Get("pwd")
	autoLogin := c.Input().Get("autoLogin") == "on"
	user, err := models.UserExistsByUserNameAndPassword(uname, pwd)
	if err != nil {
		c.FlashError("用户名或密码错误")
		c.RedirectTo(c.redirectURL)
	}
	c.SetSession("login_user_id", user.ID)
	c.SetSession("login_user_sign", utils.AuthSign(user.ID, user.Password))
	if autoLogin {
		c.Ctx.SetCookie("login_user_id", utils.ToString(user.ID), 3600*24*7)
		c.Ctx.SetCookie("login_user_sign", utils.AuthSign(user.ID, user.Password), 3600*24*7)
	}
	c.FlashSuccess("登陆成功！<br>您上次登录的时间为：" + utils.TimeDiffForHumans(models.UpdateLoginTime(user)))
	c.RedirectTo("/")
}

// @router /register [get]
func (c *UserController) Register() {
	if c.CurrentLoginUser != nil {
		c.RedirectTo("/")
	}
	c.TplName = "register.html"
	c.Data["PageTitle"] = "注册"

}

// @router /register [post]
func (c *UserController) RegisterHandler() {
	c.redirectURL = beego.URLFor("UserController.Register")
	userName := strings.TrimSpace(c.Input().Get("userName"))
	nickName := strings.TrimSpace(c.Input().Get("nickName"))
	pwd1 := c.Input().Get("pwd1")
	pwd2 := c.Input().Get("pwd2")
	if models.UserNameExists(userName) {
		c.FlashError("注册失败，用户名重复")
		c.RedirectTo(c.redirectURL)
	}
	if pwd1 != pwd2 {
		c.FlashError("注册失败，两次密码不一致")
		c.RedirectTo(c.redirectURL)
	}
	_, err := models.CreateUser(userName, pwd1, 1, nickName)
	if err != nil {
		c.FlashError("注册失败")
		c.RedirectTo(c.redirectURL)
	}
	c.FlashSuccess("注册成功")
	c.RedirectTo(beego.URLFor("UserController.Login"))
}

// // @router /password/find [get]
// func (this *UserController) FindPassword() {
// 	if this.CurrentLoginUser != nil {
// 		this.RedirectTo("/")
// 	}
// 	this.Layout = "layout/app.tpl"
// 	this.Data["PageTitle"] = "找回密码"
// 	this.Data["PageKeywords"] = this.SettingData["SEO_INDEX_KEYWORDS"]
// 	this.Data["PageDescription"] = this.SettingData["SEO_INDEX_DESCRIPTION"]
// }

// // @router /password/find [post]
// func (this *UserController) FindPasswordHandler() {
// 	this.redirectUrl = beego.URLFor("UserController.FindPassword")
// 	if this.CurrentLoginUser != nil {
// 		this.RedirectTo("/")
// 	}
// 	email := this.GetString("username")
// 	if email == "" {
// 		this.FlashError("请输入邮箱")
// 		this.RedirectTo(this.redirectUrl)
// 	}
// 	user := models.Users{}
// 	err := orm.NewOrm().QueryTable("users").Filter("email", email).One(&user)
// 	if err != nil {
// 		this.FlashError("邮箱不存在")
// 		this.RedirectTo(this.redirectUrl)
// 	}
// 	template, err := template2.ParseFiles("./views/emails/findpassword.tpl")
// 	if err != nil {
// 		this.ErrorHandler(err)
// 	}
// 	data := map[string]string{
// 		"Url": user.GenerateHashedUrl(beego.URLFor("UserController.PasswordReset")),
// 	}
// 	html := new(bytes.Buffer)
// 	err = template.Execute(html, data)
// 	if err != nil {
// 		this.ErrorHandler(err)
// 	}
// 	go utils.SendMail(email, "密码重置链接", html.String())
// 	this.FlashSuccess("密码重置邮件发送成功")
// 	this.RedirectTo(this.redirectUrl)
// }

// // @router /password/reset [post]
// func (this *UserController) PasswordResetHandler() {
// 	passwordResetData := fronted.PasswordResetValidation{}
// 	this.ValidatorAuto(&passwordResetData)

// 	userId, _ := this.GetInt("id")
// 	user, err := models.FindUserById(userId)
// 	if err != nil {
// 		this.ErrorHandler(err)
// 	}
// 	sign := this.GetString("sign")
// 	time := this.GetString("time")
// 	if user.CheckHash(sign, time) == false {
// 		this.ErrorHandler(err)
// 	}

// 	user.Password = utils.SHA256Encode(passwordResetData.Password)
// 	if _, err := orm.NewOrm().Update(user); err != nil {
// 		this.ErrorHandler(err)
// 	}
// 	this.FlashSuccess("修改成功，请重新登录")
// 	this.RedirectTo(beego.URLFor("UserController.Login"))
// }

// // @router /user/active [get]
// func (this *UserController) ActiveHandler() {
// 	userId, _ := this.GetInt("id")
// 	user, err := models.FindUserById(userId)
// 	if err != nil {
// 		this.ErrorHandler(err)
// 	}
// 	sign := this.GetString("sign")
// 	time := this.GetString("time")
// 	if user.CheckHash(sign, time) == false {
// 		this.ErrorHandler(err)
// 	}

// 	user.IsLock = models.IS_LOCK_NO
// 	if _, err := orm.NewOrm().Update(user); err != nil {
// 		this.ErrorHandler(err)
// 	}

// 	this.FlashSuccess("激活成功")
// 	this.RedirectTo("/")
// }
// @router /headimg/:id [get]
func (c *UserController) CommonHeadImg() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	user, err := models.FindUserById(id)
	if err != nil || len(user.ImgName) == 0 {
		c.RedirectTo("/static/img/akari.jpg")
	}
	imgPath := path.Join("static/uploads/headimg", user.ImgName)
	//beego.Info(imgPath)
	if !com.IsExist(imgPath) {
		c.RedirectTo("/static/img/akari.jpg")
		return
	}
	c.RedirectTo("/" + imgPath)
}
