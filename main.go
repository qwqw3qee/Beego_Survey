package main

import (
	_ "beego_survey/routers"
	"beego_survey/utils"
	"strings"

	"beego_survey/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}
func split(in string) (out []string) {
	out = strings.Split(in, "#op%")
	return
}
func addOne(in int) (out int) {
	out = in + 1
	return
}
func printResolution(record models.Records, question models.Questions, paperType int) (out string) {
	answers, _ := models.FindAnswersByQuestionIDandRecordID(question.QuestionID, record.RecordID)
	out = ""
	if len(answers) > 0 {
		out += "您的答案："
		switch question.Type {
		case 0:
			out += " " + utils.ToString(answers[0].AnswerNo)
		case 1:
			for _, v := range answers {
				out += " " + utils.ToString(v.AnswerNo)
			}
		case 2:
			out += " " + answers[0].AnswerText
		}
	}

	if paperType == 1 {
		standards, _ := models.FindStandardsByQuestion(&question)

		if len(standards) > 0 {
			out += "<br>标准答案："
			switch question.Type {
			case 0:
				out += " " + utils.ToString(standards[0].StandardNo)
			case 1:
				for _, v := range standards {
					out += " " + utils.ToString(v.StandardNo)
				}
			case 2:
				out += " " + standards[0].StandardText
			}
		}
		if len(answers) > 0 {
			out += "<br> 得分:" + utils.ToString(answers[0].Scores)
		}

	}

	return out
}
func main() {
	orm.RunSyncdb("default", false, true) //自动建表  表名  强制建表  打印相关信息

	beego.AddFuncMap("split", split)
	beego.AddFuncMap("addOne", addOne)
	beego.AddFuncMap("printResolution", printResolution)
	beego.Run()
}
