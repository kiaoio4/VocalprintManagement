package main

import (
	_ "uploadvoice/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	var FilterGateWay = func(ctx *context.Context) {
		//允许访问源
		ctx.ResponseWriter.Header().Set(
			"Access-Control-Allow-Origin",
			"*",
		)
		//允许post访问
		ctx.ResponseWriter.Header().Set(
			"Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE,UPDATE",
		)
		//header的类型
		ctx.ResponseWriter.Header().Set(
			"Access-Control-Allow-Headers",
			"Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, "+
				"Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, "+
				"X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma",
		)
		ctx.ResponseWriter.Header().Set(
			"Access-Control-Expose-Headers",
			"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control, "+
				"Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar",
		)
		ctx.ResponseWriter.Header().Set(
			"Access-Control-Max-Age",
			"1728000",
		)
		ctx.ResponseWriter.Header().Set(
			"Access-Control-Allow-Credentials",
			"true",
		)

		if ctx.Request.Method == "OPTIONS" {
			ctx.ResponseWriter.WriteHeader(200)
		}
	}

	beego.InsertFilter("*", beego.BeforeRouter, FilterGateWay)
	beego.Run()
}
