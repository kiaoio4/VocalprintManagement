package controllers

import (
	"encoding/json"
	"fmt"
	"uploadvoice/enums"
	"uploadvoice/utils"

	"github.com/astaxie/beego"
)

// JsonResult 用于返回ajax请求的基类
type JsonResult struct {
	Code enums.JsonResultCode `json:"code"`
	Msg  string               `json:"msg"`
	Data interface{}          `json:"data"`
}

//BaseController .
type BaseController struct {
	beego.Controller
	controllerName string                 // 当前控制名称
	actionName     string                 // 当前action名称
	param          map[string]interface{} // 请求参数
	dataRoot       string                 // 数据目录
}

func (c *BaseController) Prepare() {
	c.controllerName, c.actionName = c.GetControllerAndAction()
	c.jsonRequest()
	c.dataRoot = beego.AppConfig.String("dataroot")
}

/**
 * @description: 返回结果JSON
 */
func (c *BaseController) jsonResult(code enums.JsonResultCode, msg string, data interface{}) {
	r := &JsonResult{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.Data["json"] = r
	if code == 500 {
		jsonStr, _ := json.Marshal(data)
		utils.LogError(fmt.Sprintf("%s - %s - %s", string(c.Ctx.Input.RequestBody), msg, jsonStr))
	} else if code == 501 {
		jsonStr, _ := json.Marshal(data)
		utils.LogWarning(fmt.Sprintf("%s - %s - %s", string(c.Ctx.Input.RequestBody), msg, jsonStr))
	} else if code == 600 {
		jsonStr, _ := json.Marshal(data)
		utils.LogInfo(fmt.Sprintf("%s - %s - %s", string(c.Ctx.Input.RequestBody), msg, jsonStr))
	}
	c.ServeJSON()
	c.StopRun()
}

/**
 * @description: 请求参数JSON获取
 */
func (c *BaseController) jsonRequest() {
	json.Unmarshal(c.Ctx.Input.RequestBody, &c.param)
	// c.param["URI"] = fmt.Sprintf("%s[%s.%s]", c.Ctx.Request.RequestURI, c.controllerName, c.actionName)
}

/**
 * @description: 检查必选参数
 */
func (c *BaseController) checkParams(request []string) bool {
	for _, item := range request {
		if _, ok := c.param[item]; !ok {
			return false
		}
	}
	return true
}
