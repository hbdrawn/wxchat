package httpclient

import (
	"encoding/json"
	"wxchat/conf"
)

var (
	accessUrl = conf.ServerUrl + "message/send?access_token=" + conf.AccessToken
)

type textBody struct{
	Content string `json:"content"`
}

type MsgTextJson struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Text    textBody `json:"text"`
	Safe    int    `json:"safe"`
}

func SendTextMsg(content string) {
	textBody := textBody{content}
	text := &MsgTextJson{Touser: "@all", Msgtype: "text", Agentid: 0, Text: textBody, Safe: 0}
	params, _ := json.Marshal(text)
	HttpPost4Json(accessUrl, string(params))
}

func SendMsg4Warn(content string) string{
	textBody := textBody{content}
	text := &MsgTextJson{Totag:"1", Msgtype: "text", Agentid: 0, Text: textBody, Safe: 0}
	params, _ := json.Marshal(text)
	return HttpPost4Json(accessUrl, string(params))
}
