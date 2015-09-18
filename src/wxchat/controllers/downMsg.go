package controllers

import(
	"github.com/astaxie/beego"
	"wxchat/httpclient"
	
)


type DownMsg struct{
	beego.Controller
}

func (this *DownMsg) Get() {
	msg := this.GetString("msg")
	rsp := httpclient.SendMsg4Warn(msg)
	this.Ctx.WriteString(rsp)
}

func (this *DownMsg) Post() {
	msg := this.GetString("msg")
	rsp := httpclient.SendMsg4Warn(msg)
	this.Ctx.WriteString(rsp)
}