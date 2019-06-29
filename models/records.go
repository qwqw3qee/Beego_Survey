package models

import (
	"beego_survey/lib"

	"github.com/astaxie/beego/orm"
)

//Records 答题记录表
type Records struct {
	RecordID     int64   `orm:"unique;pk;auto;column(record_id)"`
	PaperID      *Papers `orm:"rel(fk);on_delete(cascade);column(paper_id)"`
	AnswerUserID *Users  `orm:"size(32);rel(fk);on_delete(cascade);column(answer_user_id)"`
	AnswerDate   lib.Time
	UseTime      int
	Scores       int
}

func CreateRecord(record *Records) (int64, error) {
	return orm.NewOrm().Insert(record)
}
func UpdateRecord(record *Records) (int64, error) {
	return orm.NewOrm().Update(record, "scores")
}
func FindRecordByPaperIDandAnswerUserID(paperID int64, answerUserID int64) (*Records, error) {
	record := new(Records)
	db := orm.NewOrm()
	err := db.QueryTable("Records").Filter("paper_id", paperID).Filter("answer_user_id", answerUserID).One(record)
	//db.LoadRelated(paper, "author_id")
	//beego.Info(paper)
	return record, err
}

func FindRecordsByAnswerUserID(answerUserID int64) ([]*Records, error) {
	records := make([]*Records, 0)
	db := orm.NewOrm()
	_, err := db.QueryTable("Records").Filter("answer_user_id", answerUserID).All(&records)
	for i := range records {
		db.LoadRelated(records[i], "paper_id")
		db.LoadRelated(records[i].PaperID, "author_id")
		db.LoadRelated(records[i], "answer_user_id")
	}
	//beego.Info("666", questions)
	return records, err
}
func CountRecordsByAnswerUserID(answerUserID int64) (int64, error) {
	db := orm.NewOrm()
	cnt, err := db.QueryTable("Records").Filter("answer_user_id", answerUserID).Count()
	//beego.Info("666", questions)
	return cnt, err
}
func FindRecordsByPaperID(paperID int64) ([]*Records, error) {
	records := make([]*Records, 0)
	db := orm.NewOrm()
	_, err := db.QueryTable("Records").Filter("paper_id", paperID).All(&records)
	for i := range records {
		db.LoadRelated(records[i], "paper_id")
		db.LoadRelated(records[i].PaperID, "author_id")
		db.LoadRelated(records[i], "answer_user_id")
	}
	//beego.Info("666", questions)
	return records, err
}
