package controllers

import (
	"beego_survey/models"
	"beego_survey/utils"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/Unknwon/com"

	"github.com/astaxie/beego/logs"
)

type UploadController struct {
	Base
}

func (c *UploadController) Prepare() {
	c.Base.Prepare()
	if c.CurrentLoginUser == nil {
		c.AjaxError("请登录", 999)
	}

}

// @router /member/upload/headimg [post]
func (c *UploadController) HeadImg() {
	file, header, err := c.GetFile("file")
	if err != nil {
		c.AjaxError("请选择头像文件", 1)
	}
	defer file.Close()
	// 文件mime判断
	mime := header.Header["Content-Type"][0]
	if mime != "image/jpeg" && mime != "image/png" && mime != "image/gif" {
		c.AjaxError("请上传有效图片文件", 2)
	}
	// 文件后缀判断
	extensions := strings.Split(header.Filename, ".")
	extension := strings.ToLower(extensions[len(extensions)-1])
	if extension != "jpg" && extension != "png" && extension != "gif" {
		c.AjaxError("请上传jpeg/png/gif图片", 3)
	}
	// 文件大小判断
	if header.Size/(1024*1024) > 2 {
		c.AjaxError("头像文件大小不能超过2MB", 4)
	}

	// 保存文件

	oldpath := path.Join("static/uploads/headimg", c.CurrentLoginUser.ImgName)
	if len(c.CurrentLoginUser.ImgName) > 0 && com.IsExist(oldpath) {
		err := os.Remove(oldpath)
		if err != nil {
			logs.Info(err)
			c.AjaxError("旧头像删除失败", 6)
		}
	}
	c.CurrentLoginUser.ImgName = ""
	newFilename := headPrefix + utils.ToString(c.CurrentLoginUser.ID) + "." + extension
	newPath := path.Join("static/uploads/headimg", newFilename)
	err = c.SaveToFile("file", newPath)
	if err != nil {
		logs.Info(err)
		c.AjaxError("头像上传失败", 5)
	}
	if _, err := models.UpdateHeadImg(c.CurrentLoginUser, newFilename); err != nil {
		c.AjaxError("更新头像数据失败", 7)
	}
	res := make(map[string]string)
	res["path"] = "/" + newPath
	c.AjaxSuccess("上传成功", res)
}

// @router /paper/upload/paperimg [post]
func (c *UploadController) PaperImg() {
	paperID, err := strconv.ParseInt(c.GetString("PaperID", ""), 10, 64)
	if paperID == 0 || err != nil {
		c.AjaxError("试卷ID不合法", 1)
	}
	var paper *models.Papers
	paper, err = models.FindPaperByPaperID(paperID)
	if err != nil {
		c.AjaxError("试卷ID不存在", 2)
	}
	file, header, err := c.GetFile("file")
	if err != nil {
		c.AjaxError("请选择图片文件", 3)
	}
	defer file.Close()
	// 文件mime判断
	mime := header.Header["Content-Type"][0]
	if mime != "image/jpeg" && mime != "image/png" && mime != "image/gif" {
		c.AjaxError("请上传有效图片文件", 4)
	}
	// 文件后缀判断
	extensions := strings.Split(header.Filename, ".")
	extension := strings.ToLower(extensions[len(extensions)-1])
	if extension != "jpg" && extension != "png" && extension != "gif" {
		c.AjaxError("请上传jpeg/png/gif图片", 5)
	}
	// 文件大小判断
	if header.Size/(1024*1024) > 2 {
		c.AjaxError("图片文件大小不能超过2MB", 6)
	}

	// 保存文件

	oldpath := path.Join("static/uploads/paperimg", paper.ImgName)
	if len(paper.ImgName) > 0 && com.IsExist(oldpath) {
		err := os.Remove(oldpath)
		if err != nil {
			logs.Info(err)
			c.AjaxError("旧图片删除失败", 7)
		}
	}
	c.CurrentLoginUser.ImgName = ""
	newFilename := paperPrefix + utils.ToString(paper.PaperID) + "." + extension
	newPath := path.Join("static/uploads/paperimg", newFilename)
	err = c.SaveToFile("file", newPath)
	if err != nil {
		logs.Info(err)
		c.AjaxError("图片上传失败", 8)
	}
	if _, err := models.UpdatePaperImg(paper, newFilename); err != nil {
		c.AjaxError("更新头像数据失败", 9)
	}
	res := make(map[string]string)
	res["path"] = "/" + newPath
	c.AjaxSuccess("上传成功", res)
}
