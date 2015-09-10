package main

import (
	_ "wxchat/routers"
	"github.com/astaxie/beego"
	"wxchat/httpclient"
)

func main() {
	
	//初始化token
	httpclient.GetToken()
	httpclient.SendTextMsg("微信企业号代理启动成功--@立联小分队@")
	
	beego.Run()
}

