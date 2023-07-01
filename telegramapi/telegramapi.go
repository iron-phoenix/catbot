package telegramapi

import "net/http"

type Chat struct {
	Id int `json:"id"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

func HandleTelegramWebHook(w http.ResponseWriter, req *http.Request) {

}
