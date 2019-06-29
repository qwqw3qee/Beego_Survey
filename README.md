# Beego_Survey

## 介绍

基于Beego + BootStrap的简易问卷调查系统

## 功能

+ [x] 账户管理
+ [x] 编辑问卷
+ [x] 做问卷
+ [x] 简易判题
+ [x] 问卷管理
+ [x] 问卷信息展示

## TODO

+ [ ] 管理员后台管理
+ [ ] 验证码
+ [ ] 统计分析
+ [ ] 导出Excel
+ [ ] 代码风格整理
+ [ ] 完善表单验证

## 依赖

| 库 | 介绍 |
| --- | --- |
| `github.com/astaxie/beego` | Beego框架 |
| `github.com/go-sql-driver/mysql` | Mysql库 |
| `github.com/Unknwon/com`|常用函数库|

## 前端页面用到的项目

|项目 | 介绍 |
| --- | --- |
|`github.com/twbs/bootstrap`|BootStrap框架|
|`github.com/jquery/jquery`|JQuery库|
|`github.com/uxsolutions/bootstrap-datepicker`|日期时间选择器|
|`github.com/wenzhixin/bootstrap-table`|BootStrap-table|
|`github.com/nghuuphuoc/bootstrapvalidator`|表单验证神器|
|`github.com/lk-code/jquery-input-counter`|数字加减输入框|
|`github.com/LiranCohen/stickUp`|stickUp|
|`github.com/CodeSeven/toastr`|toastr|

## 开发工具

* [VS Code](https://code.visualstudio.com/)

## 参考资料/项目

* [基于 Beego + Vue 开发的在线问答系统](https://github.com/Qsnh/goa)
* [网易云课堂-Go编程基础](https://study.163.com/course/courseMain.htm?courseId=306002)
* [网易云课堂-Go Web 基础](https://study.163.com/course/courseMain.htm?courseId=328001)
* [Beego开发文档](https://beego.me/docs/intro/)

## 使用说明

1. 导入数据库
    * 将 `data` 下面的SQL文件导入到数据库中
    * 默认用的是`mysql`,数据库名称为 `beego_survey`

2. 配置文件
    * 有关数据库的参数可以在`./conf/app.conf`文件中修改

``` 
    DATABASE_HOST=127.0.0.1
    DATABASE_PORT=3306
    DATABASE_USER=root
    DATABASE_PASS=123456
    DATABASE_NAME=beego_survey
    DATABASE_TIMEZONE=Asia/Shanghai
```

3. 编译运行

```bash
cd ~/gowork/src/beego_survey
go build -o beego_survey main.go
```
3. 其他说明       
数据库默认有两个账户`anonymous`和`admin`,密码均为`admin`，这两个用户暂时必须位于数据库的前两个条记录中。
