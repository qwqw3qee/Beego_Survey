package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego_survey/controllers:ApiMemberController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:ApiMemberController"],
        beego.ControllerComments{
            Method: "MemberDo",
            Router: `/api/member/do`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:ApiMemberController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:ApiMemberController"],
        beego.ControllerComments{
            Method: "MemberPublish",
            Router: `/api/member/publish`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:ApiMemberController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:ApiMemberController"],
        beego.ControllerComments{
            Method: "PublishDolist",
            Router: `/api/member/publish/dolist`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:ApiPapersController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:ApiPapersController"],
        beego.ControllerComments{
            Method: "PaperList",
            Router: `/api/paper/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:ApiUsersController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:ApiUsersController"],
        beego.ControllerComments{
            Method: "VerifyUserName",
            Router: `/register/verify`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:IndexController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:MemberController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:MemberController"],
        beego.ControllerComments{
            Method: "NickNameChangeHandler",
            Router: `/member/change_nickname`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:MemberController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:MemberController"],
        beego.ControllerComments{
            Method: "PasswordChangeHandler",
            Router: `/member/change_password`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:MemberController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:MemberController"],
        beego.ControllerComments{
            Method: "MemberDo",
            Router: `/member/do`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:MemberController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:MemberController"],
        beego.ControllerComments{
            Method: "HeadImg",
            Router: `/member/headimg`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:MemberController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:MemberController"],
        beego.ControllerComments{
            Method: "MemberIndex",
            Router: `/member/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:MemberController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:MemberController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/member/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:MemberController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:MemberController"],
        beego.ControllerComments{
            Method: "MemberPublish",
            Router: `/member/publish`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:MemberController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:MemberController"],
        beego.ControllerComments{
            Method: "PublishDolist",
            Router: `/member/publish/dolist`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:PaperController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:PaperController"],
        beego.ControllerComments{
            Method: "PaperDoHandler",
            Router: `/paper/do`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:PaperController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:PaperController"],
        beego.ControllerComments{
            Method: "PaperDo",
            Router: `/paper/do/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:PaperController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:PaperController"],
        beego.ControllerComments{
            Method: "PaperImg",
            Router: `/paper/img/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:PaperController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:PaperController"],
        beego.ControllerComments{
            Method: "PaperList",
            Router: `/paper/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:PaperController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:PaperController"],
        beego.ControllerComments{
            Method: "PaperNew",
            Router: `/paper/new`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:PaperController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:PaperController"],
        beego.ControllerComments{
            Method: "PaperNewHandler",
            Router: `/paper/new/paper`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:PaperController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:PaperController"],
        beego.ControllerComments{
            Method: "QuestionsNewHandler",
            Router: `/paper/new/questions`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:PaperController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:PaperController"],
        beego.ControllerComments{
            Method: "PaperWatch",
            Router: `/paper/watch`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:UploadController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:UploadController"],
        beego.ControllerComments{
            Method: "HeadImg",
            Router: `/member/upload/headimg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:UploadController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:UploadController"],
        beego.ControllerComments{
            Method: "PaperImg",
            Router: `/paper/upload/paperimg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:UserController"],
        beego.ControllerComments{
            Method: "CommonHeadImg",
            Router: `/headimg/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:UserController"],
        beego.ControllerComments{
            Method: "LoginHandler",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego_survey/controllers:UserController"] = append(beego.GlobalControllerRouter["beego_survey/controllers:UserController"],
        beego.ControllerComments{
            Method: "RegisterHandler",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
