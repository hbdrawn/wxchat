package httpclient

import (
	"encoding/json"
	"log"
	"sync"
	"time"
	"wxchat/conf"
)

const (
	accessTokenUrl = conf.ServerUrl + "gettoken?corpid=" + conf.WxchatCorpID + "&corpsecret=" + conf.WxchatSecret
)

var mutex sync.RWMutex

//type AccessToken struct {
//	HttpClient
//}

type tokenJson struct {
	Access_token string `json:"access_token"`
	Expires_in   int	`json:"expires_in"`
}

func  GetToken() {
	if conf.AccessToken == "" {
		mutex.RLock()
		if conf.AccessToken == "" {
			getTokenByHttp()
			go time4Token()
		}
		mutex.RUnlock()
	}
}

//定时调度，按频率获取token
func time4Token() {
	timer := time.NewTicker(conf.ExpiresIn * time.Second)
	for {
		select {
		case <-timer.C:
			getTokenByHttp()
			log.Println("定时获取AccessToken成功")
		}
	}
}

func getTokenByHttp() {
	resultStr := HttpGet(accessTokenUrl, "")
	if resultStr != "" {
		if err := ErrorHandler(resultStr); err.Errcode == 0 {
			var resultToken tokenJson
			if err := json.Unmarshal([]byte(resultStr), &resultToken); err != nil {
				log.Fatalln("获取token出错：解析json串出错")
			} else {
				conf.AccessToken = resultToken.Access_token
				conf.ExpiresIn = time.Duration(resultToken.Expires_in - 200)
				log.Printf("获取token:%s", resultToken.Access_token)
				log.Printf("获取token超时时间:%d" , resultToken.Expires_in)
			}
		}

	}else{
		log.Fatalln("获取token失败，30s后重新获取")
		conf.ExpiresIn = 30;
	}
}
