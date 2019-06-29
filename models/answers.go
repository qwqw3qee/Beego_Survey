package models

import "github.com/astaxie/beego/orm"

//Answers 答卷表
type Answers struct {
	ID         int64      `orm:"unique;pk;auto;column(id)"`
	RecordID   *Records   `orm:"rel(fk);on_delete(cascade);column(record_id)"`
	PaperID    *Papers    `orm:"rel(fk);on_delete(cascade);column(paper_id)"`
	QuestionID *Questions `orm:"rel(fk);on_delete(cascade);column(question_id)"`
	AnswerNo   int        `orm:"null;default(0)"`
	AnswerText string     `orm:"size(5000);null;default(\"\")"`
	Scores     int        `orm:"null;default(0)"`
}

func AnswersCountByPerson(userID int64) (int64, error) {
	answer := new(Answers)
	return orm.NewOrm().QueryTable(answer).Filter("record_id", userID).Count()
}
func AnswersInsertMulti(num int, answers []*Answers) error {
	db := orm.NewOrm().QueryTable("Answers")
	t, _ := db.PrepareInsert()
	for i := range answers {
		_, err := t.Insert(answers[i])
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

func FindAnswersByQuestionIDandRecordID(questionID int64, recordID int64) ([]*Answers, error) {
	answers := make([]*Answers, 0)
	db := orm.NewOrm()
	_, err := db.QueryTable("Answers").Filter("question_id", questionID).Filter("record_id", recordID).All(&answers)
	for i := range answers {
		db.LoadRelated(answers[i], "record_id")
		db.LoadRelated(answers[i], "paper_id")
		db.LoadRelated(answers[i], "question_id")
	}
	//beego.Info("666", questions)
	return answers, err
}
