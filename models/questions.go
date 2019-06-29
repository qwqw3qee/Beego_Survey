package models

import (
	"github.com/astaxie/beego/orm"
)

//Questions 问题表
type Questions struct {
	QuestionID  int64   `orm:"unique;pk;auto;column(question_id)"`
	PaperID     *Papers `orm:"rel(fk);on_delete(cascade);column(paper_id)"`
	No          int     `orm:"null;default(0)"`
	Type        int     `orm:"default(0)"`
	Description string  `orm:"size(5000);default(\"\")"`
	OptionsNum  int     `orm:"default(0)"`
	OptionsInfo string  `orm:"size(5000);null;default(\"\")"`
	Remark      string  `orm:"size(5000);default(\"\")"`
}

func QuestionsInsertMulti(num int, questions []*Questions) error {
	db := orm.NewOrm().QueryTable("Questions")
	t, _ := db.PrepareInsert()
	for i := range questions {
		_, err := t.Insert(questions[i])
		if err != nil {
			return err
		}
	}
	t.Close() // 别忘记关闭 statement
	//beego.Info(questions)
	// for i := range *questions {
	// 	db.LoadRelated(&(*questions)[i], "question_id")
	// }
	return nil
}
func QuestionsByPaperID(paperID int64) ([]*Questions, error) {
	questions := make([]*Questions, 0)
	_, err := orm.NewOrm().QueryTable("Questions").Filter("paper_id", paperID).All(&questions)
	//beego.Info("666", questions)
	return questions, err
}
func FindQuetionByQuetionID(questionID int64) (*Questions, error) {
	question := new(Questions)
	db := orm.NewOrm()
	err := db.QueryTable("Questions").Filter("question_id", questionID).One(question)
	db.LoadRelated(question, "paper_id")
	return question, err
}
