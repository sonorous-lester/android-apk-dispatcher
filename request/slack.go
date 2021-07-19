package request

import (
	"apkDispatcher/data"
	"apkDispatcher/slack"
	"log"
)

func (req *request) SendVariantBtnMessage(message slack.Selector, url string){
	_, err := req.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(message).
		Post(url)
	if err != nil {
		log.Printf("send variant button message goes wrong %v \n", err)
	}
}

func (req *request) SendBranchMenuMessage(channelId string, url string, options []data.Option){
	_, err := req.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(slack.CreateBranchMenuMsgSelector(channelId, options)).
		Post(url)
	if err != nil {
		log.Printf("send branch menu message goes wrong %v \n", err)
	}
}

func (req *request)SendMessage(url string){
	_, err := req.client.R().
		SetBody(slack.CreateMessageBlock("建立 Apk 中，請稍候")).
		Post(url)
	if err != nil {
		log.Printf("send message goes wrong %v \n", err)
	}
}