package conf

import (
	"time"
)

const(
	WxchatToken = "erutangis"
	WxchatAppId = "wx363b955d1f7d0392" //wx056e6691f0a5b45d"
	WxchatAppSecret = "735dc067d29daa35107af824b883abae"//"8efb735f41ebbabb1bc1c1cdf020c23c"
	WxchatEncodingAESKey = "242DinHMvnDazI7st4Nv53ZtSTtQ99W6muoiryxxDof"
	//user=hbdrawn@163.com hb69025087212
	//企业号信息
	WxchatCorpID = "wxd1a8c7c98084cc1c"
	WxchatSecret = "0T13yB5piVMLhW18MLd526eNuXIYgRJk1UcF7u8IOnpC0YLXh3Cs49r-yHYt2idD"
	
	ServerUrl = "https://qyapi.weixin.qq.com/cgi-bin/"
)

var(
	AccessToken = ""	//平台首次启动或定时获取
	ExpiresIn time.Duration = 10 //超时时间
)
