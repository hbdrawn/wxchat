package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"bytes"
)

const(
	MINAJSON = "application/json;charset=utf-8"
	MINAXML = "application/xml;charset=utf-8"
	MINAFORM = "application/x-www-form-urlencoded"
)

//http公共方法封装，继承此struct即可直接调用
type HttpClient struct {
}

func (httpClient *HttpClient) HttpGet(url, params string) string {
	requestUrl := []string{url, params}
	resp, err := http.Get(strings.Join(requestUrl, "?"))
	defer resp.Body.Close()
	if err != nil {
		// handle error
		fmt.Printf("http访问[%s]异常:%s", url, err.Error())
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Printf("http访问[%s]解析异常:%s", url, err.Error())
		return ""
	}

	return string(body)
}

func (httpclient *HttpClient) HttpPost(url, params string) string {

	resp, err := http.Post(url,MINAFORM , strings.NewReader(params)) //"name=cjb"

	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("http访问[%s]异常:%s", url, err.Error())
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		fmt.Printf("http访问[%s]解析异常:%s", url, err.Error())
		return ""
	}

	return string(body)
}


func (httpclient *HttpClient) HttpPost4Json(url, params string) string {

	resp, err := http.Post(url, MINAJSON, bytes.NewBuffer([]byte(params))) //"name=cjb"

	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("http访问[%s]异常:%s", url, err.Error())
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		fmt.Printf("http访问[%s]解析异常:%s", url, err.Error())
		return ""
	}

	return string(body)
}


func (httpclient *HttpClient) HttpPost4Xml(url, params string) string {

	resp, err := http.Post(url, MINAXML, bytes.NewBuffer([]byte(params)))

	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("http访问[%s]异常:%s", url, err.Error())
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
		fmt.Printf("http访问[%s]解析异常:%s", url, err.Error())
		return ""
	}

	return string(body)
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
