package routers

import (
	"beego_survey/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.MemberController{})
	beego.Include(&controllers.PaperController{})
	beego.Include(&controllers.UploadController{})
	beego.Include(&controllers.ApiUsersController{})
	beego.Include(&controllers.ApiPapersController{})
	beego.Include(&controllers.ApiMemberController{})

	//beego.SetStaticPath("/static/uploads/headimg", "/headimg")
}
