package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("username", "", -1, "/")
		this.Ctx.SetCookie("password", "", -1, "/") //-1 立即删除效果

		this.Redirect("/", 301)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	autoLogin := this.Input().Get("autoLogin") == "on"
	if beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		this.Ctx.SetCookie("username", username, maxAge, "/")
		this.Ctx.SetCookie("password", password, maxAge, "/")
	}

	this.Redirect("/", 301)
	return
}

func CheckAccount(this *context.Context) bool {
	ck, err := this.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value

	ck, err = this.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := ck.Value

	return beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password
}
