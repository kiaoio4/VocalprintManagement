package controllers

import (
	"fmt"
	"os"
	"path"
	"uploadvoice/enums"
	"uploadvoice/utils"
)

//UserController .
type UserController struct {
	BaseController
}

//VoiceUpload .
func (c *UserController) VoiceUpload() {

	f, h, _ := c.GetFile("file")
	username := c.Ctx.Input.Query("User")
	id := c.Ctx.Input.Query("Id")

	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".wav": true,
		".mp3": true,
		".md":  true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		c.jsonResult(enums.JRCodeRequestError, "后缀名不符合,上传文件失败，请重新上传", "")
		return
	}

	filedir := c.dataRoot + "/" + username + "/" + id
	if !utils.IsExist(filedir) {
		if err := os.MkdirAll(filedir, 0777); err != nil {
			c.jsonResult(enums.JRCodeRequestError, fmt.Sprintf("%v", err), nil)
		}
	}
	fpath := filedir + "/" + h.Filename
	defer f.Close()
	err := c.SaveToFile("file", fpath)
	if err != nil {
		c.jsonResult(enums.JRCodeRequestError, fmt.Sprintf("%v", err), nil)
	}

	c.jsonResult(enums.JRCodeSucc, "上传文件成功", map[string]interface{}{
		"Id":       id,
		"Name":     username,
		"FileName": h.Filename,
	})
}
