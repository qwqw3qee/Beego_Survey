package models

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	//导入go-sqlite3库
	//_ "github.com/mattn/go-sqlite3"
	//"github.com/Unknwon/com"

	//导入mysql库
	_ "github.com/go-sql-driver/mysql"
)

const (
	//dbName        = "data/beego_survey.db" //_DB_NAME 数据库名称
	//sqlite3Driver = "sqlite3"              //_SQLITE3_DRIVER 驱动名称
	mysqlDriver = "mysql"
)

//RegisterDB 注册数据库，在main函数中调用
func RegisterDB() {
	/**sqlite3
	if !com.IsExist(dbName) {
		os.MkdirAll(path.Dir(dbName), os.ModePerm)
		os.Create(dbName)
	}
	*/
	orm.RegisterModel(new(Papers), new(Users), new(Questions), new(Records), new(Answers), new(Standards)) //注册模型
	//mysql
	dbHost := beego.AppConfig.String("DATABASE_HOST")
	dbPort := beego.AppConfig.String("DATABASE_PORT")
	dbUser := beego.AppConfig.String("DATABASE_USER")
	dbPass := beego.AppConfig.String("DATABASE_PASS")
	dbName := beego.AppConfig.String("DATABASE_NAME")
	dbTimeZone := beego.AppConfig.String("DATABASE_TIMEZONE")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4"
	if dbTimeZone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(dbTimeZone)
	}
	orm.RegisterDriver(mysqlDriver, orm.DRMySQL)
	orm.RegisterDataBase("default", mysqlDriver, dsn, 30)
	//以下是sqlite3
	//orm.RegisterDriver(sqlite3Driver, orm.DRSqlite)//注册驱动 Sqlite默认会注册
	//orm.RegisterDataBase("default", sqlite3Driver, dbName+"?_foreign_keys=true", 10) //创建默认数据库 必须有一个defult  最大连接数
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true //调试模式
	}
}
