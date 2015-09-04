package controllers

import (
	"github.com/astaxie/beego"
	"wxchat/utils"
	"strings"
	"fmt"
	"crypto/sha1"
	"io"
	"wxchat/conf"
)

type SignatureContoller struct {
	beego.Controller
}

func (this *SignatureContoller) Get() {
	signature := strings.ToLower(this.GetString("signature"))
	timestamp := strings.ToLower(this.GetString("timestamp"))
	nonce := strings.ToLower(this.GetString("nonce"))
	echostr := this.GetString("echostr")

	//1. 将token、timestamp、nonce三个参数进行字典序排序
	//2. 将三个参数字符串拼接成一个字符串进行sha1加密
	//3. 开发者获得加密后的字符串可与signature对比，标识该请求来源于微信
	result := sortParams(conf.WxchatToken,timestamp,nonce)	
	fmt.Println(">>>%s",result)
	t := sha1.New()
	io.WriteString(t,result)
	strTmp := fmt.Sprintf("%x",t.Sum(nil))
	if(strTmp == signature){
		fmt.Println("验证成功，即将返回微信服务器")
		this.Ctx.WriteString(echostr)
	}else{
		fmt.Println(">>>验证失败:timestamp:%s,nonce:%s,signatrue:%s",timestamp,nonce,signature)
		this.Abort("error")
	}
}

func sortParams(token, timestamp, nonce string) string {
	tokens := make([]string, 3, 3)
	//先对token和timestamp进行排序
	i := utils.CompareStr(token,timestamp)
	j := utils.CompareStr(token,nonce)
	k := utils.CompareStr(timestamp,nonce)
	if i == 1 && j == -1{
		//token > timestamp  && nonce > token
		tokens[0] = timestamp
		tokens[1] = token
		tokens[2] = nonce
	}else if (i == -1 && k == -1){
		//token < timestamp < nonce
		tokens[1] = timestamp
		tokens[0] = token
		tokens[2] = nonce
	}else if ( j== 1 && k ==-1){
		tokens[0] = timestamp
		tokens[2] = token
		tokens[1] = nonce
	}else if (j == -1 && k == 1){
		tokens[2] = timestamp
		tokens[0] = token
		tokens[1] = nonce
	}else if (k == 1 && i == 1){
		tokens[1] = timestamp
		tokens[2] = token
		tokens[0] = nonce
	}else {
		tokens[2] = timestamp
		tokens[1] = token
		tokens[0] = nonce
	}

	return strings.Join(tokens,"")
}
