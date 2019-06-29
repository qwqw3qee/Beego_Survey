package controllers

import (
	"beego_survey/jsonformat"
	"beego_survey/lib"
	"beego_survey/models"
	"beego_survey/utils"
	"encoding/json"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/Unknwon/com"
)

type PaperController struct {
	Base
}

func checkLogin(c *PaperController) {
	if c.CurrentLoginUser == nil {
		c.RedirectTo(beego.URLFor("UserController.Login"))
	}
}

// @router /paper/img/:id [get]
func (c *PaperController) PaperImg() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	paper, err := models.FindPaperByPaperID(id)
	if err != nil {
		c.RedirectTo("/static/img/paper.jpg")
	}
	if len(paper.ImgName) == 0 {
		c.RedirectTo("/static/img/paper.jpg")
	}
	imgPath := path.Join("static/uploads/paperimg", paper.ImgName)
	//beego.Info(imgPath)
	if !com.IsExist(imgPath) {
		c.RedirectTo("/static/img/paper.jpg")
		return
	}
	c.RedirectTo("/" + imgPath)
}

// @router /paper/list [get]
func (c *PaperController) PaperList() {
	c.TplName = "paper_list.html"
	c.Data["IsPaperList"] = true
	c.Data["PageTitle"] = "问卷列表"
}

// @router /paper/new [get]
func (c *PaperController) PaperNew() {
	checkLogin(c)
	c.TplName = "paper_new.html"
	c.Data["IsPaperNew"] = true
	c.Data["PageTitle"] = "发布问卷"
}

// @router /paper/new/paper [post]
func (c *PaperController) PaperNewHandler() {
	checkLogin(c)
	var paperJson jsonformat.PaperJson
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &paperJson)
	if err != nil {
		c.AjaxError("解析json出错", 1)
	}
	tempType, _ := strconv.Atoi(paperJson.Type)
	loc, _ := time.LoadLocation("Local")
	tempStartTime, _ := time.ParseInLocation("2006-01-02 15:04:05", paperJson.StartDate, loc)
	tempEndTime, _ := time.ParseInLocation("2006-01-02 15:04:05", paperJson.EndDate, loc)
	tempTimeLimit, _ := strconv.Atoi(paperJson.TimeLimit)
	paper := models.Papers{
		AuthorID:    c.CurrentLoginUser,
		Type:        tempType,
		Title:       paperJson.Title,
		Subtitle:    paperJson.Subtitle,
		Description: paperJson.Description,
		AddDate:     lib.Time{Time: time.Now()},
		StartDate:   lib.Time{Time: tempStartTime},
		EndDate:     lib.Time{Time: tempEndTime},
		TimeLimit:   tempTimeLimit,
		Remark:      paperJson.Remark,
	}
	beego.Info(paper)
	if _, err := models.CreatePaper(&paper); err != nil {
		c.AjaxError("插入数据库出错", 2)
	}
	c.AjaxSuccess("试卷创建成功", paper.PaperID)
}

// @router /paper/new/questions [post]
func (c *PaperController) QuestionsNewHandler() {
	checkLogin(c)
	type DataJson struct {
		PaperID int64
		Rows    []jsonformat.QuestionsJson `json:"rows"`
		Total   int                        `json:"total"`
	}
	var questionsJson DataJson
	data := c.Ctx.Input.RequestBody
	// str := string(data[:])
	// beego.Info(str)
	err := json.Unmarshal(data, &questionsJson)
	if err != nil {
		c.AjaxError("解析json出错", 1)
	}
	var paper *models.Papers
	paper, err = models.FindPaperByPaperID(questionsJson.PaperID)
	if err != nil {
		c.AjaxError("查找问卷ID失败", 1)
	}
	hasAns := paper.Type == 1
	questions := make([]*models.Questions, questionsJson.Total)
	standards := make([]*models.Standards, 0)
	for i, v := range questionsJson.Rows {
		questions[i] = &models.Questions{
			PaperID:     paper,
			No:          utils.String2Int(v.No),
			Type:        utils.String2Int(v.Type),
			Description: v.Description,
			OptionsNum:  utils.String2Int(v.OptionsNum),
			OptionsInfo: v.OptionsInfo,
			Remark:      v.Remark,
		}
	}
	if nil != models.QuestionsInsertMulti(questionsJson.Total, questions) {
		c.AjaxError("插入题目失败", 1)
	}
	responseData := map[string]interface{}{}
	if hasAns {
		cnt := 0
		for i, v := range questionsJson.Rows {
			switch questions[i].Type {
			case 0:
				standards = append(standards, &models.Standards{
					PaperID:      paper,
					QuestionID:   questions[i],
					StandardNo:   utils.String2Int(v.StandardNo),
					StandardText: "",
					Scores:       utils.String2Int(v.Scores),
				})
				cnt++
			case 1:
				list := strings.Split(v.StandardNo, "#op%")
				beego.Info(list)
				for _, s := range list {
					standards = append(standards, &models.Standards{
						PaperID:      paper,
						QuestionID:   questions[i],
						StandardNo:   utils.String2Int(s),
						StandardText: "",
						Scores:       utils.String2Int(v.Scores),
					})
					cnt++
				}
			case 2:
				standards = append(standards, &models.Standards{
					PaperID:      paper,
					QuestionID:   questions[i],
					StandardNo:   0,
					StandardText: v.StandardText,
					Scores:       utils.String2Int(v.Scores),
				})
				cnt++
			}

		}
		if nil != models.StandardsInsertMulti(cnt, standards) {
			c.AjaxError("插入题目失败", 1)
		}
		responseData["standards"] = cnt
	}

	responseData["questions"] = questionsJson.Total
	c.Data["json"] = responseData
	c.AjaxSuccess("题目创建成功", responseData)
}

// @router /paper/do/:id [get]
func (c *PaperController) PaperDo() {
	//checkLogin(c)
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	paper, err := models.FindPaperByPaperID(id)
	if err != nil {
		c.FlashError("问卷id不合法")
		c.RedirectTo("/")
	}
	var questions []*models.Questions
	questions, err = models.QuestionsByPaperID(paper.PaperID)
	if err != nil {
		c.FlashError("提取题目出错")
		c.RedirectTo("/")
	}
	//beego.Info(paper)
	//beego.Info(questions)
	c.TplName = "paper_do.html"
	c.Data["PageTitle"] = "正在答题"
	c.Data["paper"] = paper
	c.Data["questions"] = questions
}

// @router /paper/do [post]
func (c *PaperController) PaperDoHandler() {
	//checkLogin(c)
	type DataJson struct {
		PaperID      int64
		AnswerUserID int64
		UseTime      int
		AnswerDate   int64
		Rows         []jsonformat.AnswersJson `json:"rows"`
		Total        int                      `json:"total"`
	}
	var answersJson DataJson
	data := c.Ctx.Input.RequestBody
	// str := string(data[:])
	// beego.Info(str)
	err := json.Unmarshal(data, &answersJson)
	if err != nil {
		c.AjaxError("解析json出错", 1)
	}
	beego.Info(answersJson)
	var paper *models.Papers
	paper, err = models.FindPaperByPaperID(answersJson.PaperID)
	if err != nil {
		c.AjaxError("查找问卷ID失败", 2)
	}
	var answerUser *models.Users
	if c.CurrentLoginUser == nil {
		answerUser, _ = models.FindUserById(answersJson.AnswerUserID)
	} else {
		answerUser = c.CurrentLoginUser
	}

	record := &models.Records{
		PaperID:      paper,
		AnswerUserID: answerUser,
		AnswerDate:   lib.Time{Time: time.Unix(answersJson.AnswerDate/1000, 0)},
		UseTime:      answersJson.UseTime,
	}

	if _, err = models.CreateRecord(record); err != nil {
		c.AjaxError("创建答题记录失败", 3)
	}

	answers := make([]*models.Answers, 0)

	hasAns := paper.Type == 1
	cnt := 0
	scores := 0
	for _, v := range answersJson.Rows {
		t_question, _ := models.FindQuetionByQuetionID(utils.String2Int64(v.QuestionID))
		switch t_question.Type {
		case 0:
			t_answer := &models.Answers{
				RecordID:   record,
				PaperID:    paper,
				QuestionID: t_question,
				AnswerNo:   utils.String2Int(v.AnswerNo),
				AnswerText: "",
				Scores:     0,
			}
			if hasAns {
				t_stand, _ := models.FindStandardsByQuestion(t_question)
				if t_stand[0].StandardNo == t_answer.AnswerNo {
					t_answer.Scores = t_stand[0].Scores
					scores += t_stand[0].Scores
				}
			}
			answers = append(answers, t_answer)
			cnt++
		case 1:
			list := strings.Split(v.AnswerNo, "#op%")
			var t_stand []*models.Standards
			isWrong := false
			standardMap := map[int]bool{}
			if hasAns {
				t_stand, _ = models.FindStandardsByQuestion(t_question)
				if len(t_stand) != len(list) {
					isWrong = true
				}
				for _, b := range t_stand {
					standardMap[b.StandardNo] = true
				}
			}
			t_answer := &models.Answers{
				RecordID:   record,
				PaperID:    paper,
				QuestionID: t_question,
				AnswerNo:   0,
				AnswerText: "",
				Scores:     0,
			}
			for j, b := range list {
				no := utils.String2Int(b)
				tt_answer := *t_answer
				tt_answer.AnswerNo = no
				if hasAns {
					if isWrong || !standardMap[no] {
						isWrong = true
					}
					if !isWrong && j == len(list)-1 {
						tt_answer.Scores = t_stand[0].Scores
						scores += t_stand[0].Scores
					}
				}
				answers = append(answers, &tt_answer)
				cnt++
			}
		case 2:
			t_answer := &models.Answers{
				RecordID:   record,
				PaperID:    paper,
				QuestionID: t_question,
				AnswerNo:   0,
				AnswerText: v.AnswerText,
				Scores:     0,
			}
			if hasAns {
				t_stand, _ := models.FindStandardsByQuestion(t_question)
				if t_stand[0].StandardText == t_answer.AnswerText {
					t_answer.Scores = t_stand[0].Scores
					scores += t_stand[0].Scores
				}
			}
			answers = append(answers, t_answer)
			cnt++
		}
	}

	if hasAns {
		record.Scores = scores
		models.UpdateRecord(record)
	}
	if models.AnswersInsertMulti(cnt, answers) != nil {
		c.AjaxError("创建答题题目失败", 4)
	}
	responseData := map[string]interface{}{}
	responseData["PaperID"] = paper.PaperID
	responseData["AnswerUserID"] = record.AnswerUserID.ID
	if hasAns {
		responseData["Scores"] = record.Scores
		c.AjaxSuccess("答题完成，您的分数:"+utils.ToString(record.Scores)+"分，点确定查看试卷", responseData)
	}
	c.AjaxSuccess("答题完成，点确定查看试卷", responseData)
}

// @router /paper/watch [get]
func (c *PaperController) PaperWatch() {
	checkLogin(c)
	c.redirectURL = beego.URLFor("PapersController.PaperList")
	paperID, err := strconv.ParseInt(c.Input().Get("paperid"), 10, 64)
	if err != nil {
		c.FlashError("试卷ID参数不正确")
		c.RedirectTo(c.redirectURL)
	}
	var answerUserID int64
	answerUserID, err = strconv.ParseInt(c.Input().Get("answeruserid"), 10, 64)
	if err != nil {
		c.FlashError("用户ID参数不正确")
		c.RedirectTo(c.redirectURL)
	}
	var record *models.Records
	record, err = models.FindRecordByPaperIDandAnswerUserID(paperID, answerUserID)
	if err != nil {
		c.FlashError("未查到答题记录")
		c.RedirectTo(c.redirectURL)
	}
	var paper *models.Papers
	paper, err = models.FindPaperByPaperID(paperID)
	if err != nil {
		c.FlashError("未查到试卷")
		c.RedirectTo(c.redirectURL)
	}
	if answerUserID != c.CurrentLoginUser.ID && answerUserID != paper.AuthorID.ID && c.CurrentLoginUser.UserType != 2 {
		c.FlashError("您没有权限查看此问卷")
		c.RedirectTo(c.redirectURL)
	}
	questions, _ := models.QuestionsByPaperID(paperID)
	c.TplName = "paper_watch.html"
	c.Data["PageTitle"] = "查看问卷"
	c.Data["paper"] = paper
	c.Data["questions"] = questions
	c.Data["record"] = record

}
