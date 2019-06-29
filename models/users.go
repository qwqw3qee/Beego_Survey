package models

import (
	"beego_survey/lib"
	"beego_survey/utils"
	"time"

	"github.com/astaxie/beego/orm"
)

//Users 用户信息表
type Users struct {
	ID            int64    `orm:"unique;pk;auto;column(id)"`
	UserName      string   `orm:"size(32);unique" json:"-"`
	NickName      string   `orm:"size(32);default(\"\")"`
	UserType      int      `orm:"default(1)" json:"-"`
	Password      string   `orm:"size(256);default(\"\")" json:"-"`
	ImgName       string   `orm:"size(256);null;default(\"\")" json:"-"`
	RegisterTime  lib.Time `json:"-"`
	LastLoginTime lib.Time `json:"-"`
}

func FindUserById(id int64) (*Users, error) {
	user := new(Users)
	db := orm.NewOrm()
	err := db.QueryTable(user).Filter("id", id).One(user)
	return user, err
}

//UserExistsByEmailAndPassword 判断用户名密码是否匹配
func UserExistsByUserNameAndPassword(userName string, password string) (*Users, error) {
	//beego.Info(userName)
	//beego.Info(utils.SHA256Encode(password))
	user := new(Users)
	db := orm.NewOrm()
	err := db.QueryTable(user).Filter("user_name", userName).Filter("password", utils.SHA256Encode(password)).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func UpdateLoginTime(user *Users) lib.Time {
	lastTime := user.LastLoginTime
	user.LastLoginTime = lib.Time{Time: time.Now()}
	orm.NewOrm().Update(user, "last_login_time")
	return lastTime
}
func UpdateHeadImg(user *Users, imgName string) (int64, error) {
	user.ImgName = imgName
	return orm.NewOrm().Update(user, "img_name")
}
func CreateUser(userName string, password string, userType int, nickName string) (int64, error) {
	user := &Users{
		UserName:      userName,
		Password:      utils.SHA256Encode(password),
		UserType:      userType,
		NickName:      nickName,
		RegisterTime:  lib.Time{Time: time.Now()},
		LastLoginTime: lib.Time{Time: time.Now()},
	}
	//beego.Info(user)
	return orm.NewOrm().Insert(user)
}
func UserNameExists(name string) bool {
	user := new(Users)
	return orm.NewOrm().QueryTable(user).Filter("user_name", name).One(user) == nil
}

func PasswordUpdate(user *Users, password string) (err error) {
	user.Password = utils.SHA256Encode(password)
	_, err = orm.NewOrm().Update(user, "password")
	return err
}

func NickNameUpdate(user *Users, nickName string) (err error) {
	user.NickName = nickName
	_, err = orm.NewOrm().Update(user, "nick_name")
	return err
}
