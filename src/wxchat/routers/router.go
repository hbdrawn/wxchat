package routers

import (
	"github.com/astaxie/beego"
	"wxchat/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/signature", &controllers.SignatureContoller{})
}
