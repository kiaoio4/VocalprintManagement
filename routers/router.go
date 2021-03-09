package routers

import (
	"uploadvoice/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//初始化 namespace 对象
	ns :=
		beego.NewNamespace("/api",
		    //支持满足条件的就执行该 namespace
			beego.NSCond(func(ctx *context.Context) bool {
				return true
			}),
			//嵌套namespace
			beego.NSNamespace("userinfo",
				beego.NSRouter("upload", &controllers.UserController{}, "Post:VoiceUpload"), //upload voice file
			),
		)
	beego.AddNamespace(ns)
}
