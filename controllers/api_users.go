package controllers

import (
	"beego_survey/models"

	"strconv"
)

type ApiUsersController struct {
	Base
}

// @router /register/verify [post]
func (c *ApiUsersController) VerifyUserName() {
	toVerify := c.Input().Get("userName")
	isOk := !models.UserNameExists(toVerify)
	data := map[string]string{"valid": strconv.FormatBool(isOk)}
	c.Data["json"] = data
	c.ServeJSON()
	c.StopRun()
}
