package routers

import (
	"github.com/astaxie/beego"
	"wxchat/controllers"
)

func init() {
	beego.Router("/wxchat", &controllers.MainController{})
	beego.Router("/wxchat/signature", &controllers.SignatureContoller{})
	
	beego.Router("/wxchat/send", &controllers.DownMsg{})
}
