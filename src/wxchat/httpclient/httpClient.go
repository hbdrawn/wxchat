package httpclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"wxchat/conf"
)

const (
	MINAJSON = "application/json;charset=utf-8"
	MINAXML  = "application/xml;charset=utf-8"
	MINAFORM = "application/x-www-form-urlencoded"
)

func HttpGet(url, params string) string {
	log.Printf("http请求[%s]发送信息：\n%s", url, params)
	var requestUrl string
	if params == "" {
		requestUrl = url
	} else {
		strs := []string{url, params}
		requestUrl = strings.Join(strs, "?")
	}

	resp, err := http.Get(requestUrl)

	if err != nil {
		// handle error
		log.Printf("http访问[%s]异常:%s", url, err.Error())
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Printf("http访问[%s]解析异常:%s", url, err.Error())
		return ""
	}

	log.Printf("http请求[%s]返回值：\n%s", url, string(body))
	return string(body)
}

func HttpPost(url, params string) string {

	resp, err := http.Post(url, MINAFORM, strings.NewReader(params)) //"name=cjb"

	if err != nil {
		log.Printf("http访问[%s]异常:%s", url, err.Error())
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		log.Printf("http访问[%s]解析异常:%s", url, err.Error())
		return ""
	}

	return string(body)
}

func HttpPost4Json(url, params string) string {
	url = url + conf.AccessToken
	log.Printf("http请求[%s]发送信息：\n%s", url, params)
	resp, err := http.Post(url, MINAJSON, bytes.NewBuffer([]byte(params)))

	if err != nil {
		log.Printf("http请求[%s]异常:\n%s", url, err.Error())
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		log.Printf("http请求[%s]解析异常:\n%s", url, err.Error())
		return ""
	}
	log.Printf("http请求[%s]返回值：\n%s", url, string(body))
	return string(body)
}

func HttpPost4Xml(url, params string) string {

	resp, err := http.Post(url, MINAXML, bytes.NewBuffer([]byte(params)))

	if err != nil {
		log.Printf("http访问[%s]异常:%s", url, err.Error())
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		log.Printf("http访问[%s]解析异常:%s", url, err.Error())
		return ""
	}

	return string(body)
}

//微信系统返回的统一错误信息结构体
type ErrCode struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//统一异常处理
func ErrorHandler(body string) ErrCode {
	if strings.Contains(body, "errcode") {
		var err ErrCode
		if err1 := json.Unmarshal([]byte(body), &err); err1 == nil {
			log.Fatalf("错误码解析出错：%s", err1)
		}
		log.Fatalf("远程服务返回异常信息:%d:%s", err.Errcode, err.Errmsg)
		return err
	}
	return ErrCode{0, ""}
}

//func httpPostForm() {
//    resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
//        url.Values{"key": {"Value"}, "id": {"123"}})

//    if err != nil {
//        // handle error
//    }

//    defer resp.Body.Close()
//    body, err := ioutil.ReadAll(resp.Body)
//    if err != nil {
//        // handle error
//    }

//    fmt.Println(string(body))

//}

////同上面的post请求，必须要设定Content-Type为application/x-www-form-urlencoded，post参数才可正常传递。
//func httpDo() {
//    client := &http.Client{}

//    req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
//    if err != nil {
//        // handle error
//    }

//    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//    req.Header.Set("Cookie", "name=anny")

//    resp, err := client.Do(req)

//    defer resp.Body.Close()

//    body, err := ioutil.ReadAll(resp.Body)
//    if err != nil {
//        // handle error
//    }

//    fmt.Println(string(body))
//}
