package telegramapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

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

const (
	telegralBotAPIURL = "https://api.telegram.org/bot"

	SendMessageMethod = "sendMessage"
	SendPhotoMethod   = "sendPhoto"
)

type DataGetter interface {
	GetData(msg string) (string, string, error)
}

func SendDataToTelegramChat(dataGetter DataGetter, msg string, chatId int, method string) (string, error) {
	log.Printf("Sending data to chat %d", chatId)

	requestUrl, err := url.JoinPath(telegralBotAPIURL+os.Getenv("TELEGRAM_TOKEN"), method)
	if err != nil {
		log.Fatal("Can't get Telegram Bot Api request string", err.Error())
		return "", err
	}

	key, data, err := dataGetter.GetData(msg)
	if err != nil {
		log.Fatal("Can't get data", err.Error())
		return "", err
	}

	values := url.Values{}
	values.Add("chat_id", strconv.Itoa(chatId))
	values.Add(key, data)

	resp, err := http.PostForm(requestUrl, values)

	if err != nil {
		log.Fatal("Can't send request to Telegram Bot API", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Can't parse response from Telagram Bot API", err.Error())
		return "", err
	}

	return string(bodyBytes), nil
}

func ParseUpdate(req *http.Request) (*Update, error) {
	var update Update
	err := json.NewDecoder(req.Body).Decode(&update)
	return &update, err
}
