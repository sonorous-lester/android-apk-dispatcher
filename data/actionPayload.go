package data

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/url"
)

type ActionPayload struct {
	CallbackId  string   `json:"callback_id"`
	ResponseUrl string   `json:"response_url"`
	Channel     Channel  `json:"channel"`
	User        User     `json:"user"`
	Actions     []Action `json:"actions"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Action struct {
	Name            string     `json:"name"`
	Text            string     `json:"text"`
	Type            string     `json:"type"`
	Value           string     `json:"value"`
	SelectedOptions []Selected `json:"selected_options"`
	Options         []Option   `json:"options"`
}

type Selected struct {
	Value string `json:"value"`
}

type Option struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

func ConvertBodyToActionPayload(r io.Reader) *ActionPayload{
	body, _ := ioutil.ReadAll(r)
	strBody := string(body)
	decodeBody, _ := url.QueryUnescape(strBody)
	subStr := decodeBody[8:]
	var actionPayload ActionPayload
	err := json.Unmarshal([]byte(subStr), &actionPayload)
	if err != nil {
		log.Println("json actionPayload goes wrong", err)
		return nil
	}
	return &actionPayload
}
