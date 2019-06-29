package models

import (
	"github.com/astaxie/beego/orm"
)

//Standards 标准答案表
type Standards struct {
	ID           int64      `orm:"unique;pk;auto;column(id)"`
	PaperID      *Papers    `orm:"rel(fk);on_delete(cascade);column(paper_id)"`
	QuestionID   *Questions `orm:"rel(fk);on_delete(cascade);column(question_id)"`
	StandardNo   int        `orm:"null;default(0)"`
	StandardText string     `orm:"size(5000);null;default(\"\")"`
	Scores       int        `orm:"default(0)"`
}

func StandardsInsertMulti(num int, standards []*Standards) error {
	db := orm.NewOrm().QueryTable("Standards")
	t, _ := db.PrepareInsert()
	for i := range standards {
		_, err := t.Insert(standards[i])
		if err != nil {
			return err
		}
	}
	t.Close() // 别忘记关闭 statement
	return nil
}
func FindStandardsByPaper(paper *Papers) ([]*Standards, error) {
	standards := make([]*Standards, 0)
	db := orm.NewOrm()
	_, err := db.QueryTable("Standards").Filter("paper_id", paper).All(&standards)
	for i, _ := range standards {
		db.LoadRelated(standards[i], "question_id")
	}
	return standards, err
}
func FindStandardsByQuestion(question *Questions) ([]*Standards, error) {
	standards := make([]*Standards, 0)
	db := orm.NewOrm()
	_, err := db.QueryTable("Standards").Filter("question_id", question).All(&standards)
	return standards, err
}
