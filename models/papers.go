package models

import (
	"beego_survey/lib"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
)

//Papers 问卷表
type Papers struct {
	PaperID     int64  `orm:"unique;pk;auto;column(paper_id)"`
	AuthorID    *Users `orm:"rel(fk);on_delete(cascade);column(author_id)"`
	Type        int    `orm:"default(0)"`
	Title       string `orm:"size(256);null;default(\"\")"`
	Subtitle    string `orm:"size(256);null;default(\"\")"`
	Description string `orm:"size(5000);null;default(\"\")"`
	AddDate     lib.Time
	StartDate   lib.Time
	EndDate     lib.Time
	TimeLimit   int    `orm:"default(0)"`
	Remark      string `orm:"size(5000);default(\"\")"`
	ImgName     string `orm:"size(256);null;default(\"\")"`
}

func PapersCountByPerson(userID int64) (int64, error) {
	paper := new(Papers)
	return orm.NewOrm().QueryTable(paper).Filter("author_id", userID).Count()
}
func FindPaperByAuthorID(authorID int64) ([]*Papers, error) {
	papers := make([]*Papers, 0)
	db := orm.NewOrm()
	_, err := db.QueryTable("Papers").Filter("author_id", authorID).All(&papers)
	for i := range papers {
		db.LoadRelated(papers[i], "author_id")
	}

	//beego.Info(paper)
	return papers, err
}
func FindPaperByPaperID(paperID int64) (*Papers, error) {
	paper := new(Papers)
	db := orm.NewOrm()
	err := db.QueryTable(paper).Filter("paperID", paperID).One(paper)
	db.LoadRelated(paper, "author_id")
	//beego.Info(paper)
	return paper, err
}

func PapersAll(paper *[]Papers) (int64, error) {
	db := orm.NewOrm()
	cols, err := db.QueryTable("papers").All(paper)
	for i := range *paper {
		db.LoadRelated(&(*paper)[i], "author_id")
	}
	beego.Info(&paper)
	return cols, err
}
func TableColCount(table string) (int64, error) {
	return orm.NewOrm().QueryTable(table).Count()
}

func PapersPart(paper *[]Papers, limit int64, offset int64, sortCol string, desc bool) (cols int64, err error) {
	cols, _ = TableColCount("Papers")
	db := orm.NewOrm()
	if len(sortCol) == 0 {
		_, err = db.QueryTable("papers").Limit(limit, offset).All(paper)
	} else {
		if desc {
			sortCol = "-" + sortCol
		}
		_, err = db.QueryTable("papers").OrderBy(sortCol).Limit(limit, offset).All(paper)
	}

	for i := range *paper {
		db.LoadRelated(&(*paper)[i], "author_id")
	}
	//beego.Info(&paper)
	return
}
func CreatePaper(paper *Papers) (int64, error) {
	return orm.NewOrm().Insert(paper)
}
func UpdatePaperImg(paper *Papers, imgName string) (int64, error) {
	paper.ImgName = imgName
	return orm.NewOrm().Update(paper, "img_name")
}
