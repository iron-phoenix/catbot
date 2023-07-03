package main

import (
	"catbot/telegramapi"
	"catbot/thecatapi"
	"log"
	"net/http"
	"os"
	"strings"
)

var catVocabulary = map[string]bool{
	"кот":         true,
	"кота":        true,
	"коту":        true,
	"коте":        true,
	"коты":        true,
	"котэ":        true,
	"котя":        true,
	"котю":        true,
	"коти":        true,
	"котов":       true,
	"котейка":     true,
	"котейку":     true,
	"котейке":     true,
	"котейки":     true,
	"котеек":      true,
	"котька":      true,
	"котьку":      true,
	"котьке":      true,
	"котьки":      true,
	"котек":       true,
	"котан":       true,
	"котана":      true,
	"котану":      true,
	"котане":      true,
	"котаны":      true,
	"котанов":     true,
	"котеечка":    true,
	"котеечку":    true,
	"котеечке":    true,
	"котеечки":    true,
	"котеечек":    true,
	"котейшество": true,
	"котейшеств":  true,
	"котейшества": true,
	"котейшеству": true,
	"котейшестве": true,
	"котей":       true,
	"котеи":       true,
	"котея":       true,
	"котею":       true,
	"котофей":     true,
	"котофея":     true,
	"котофею":     true,
	"котофеи":     true,
	"котофейка":   true,
	"котофейку":   true,
	"котофейке":   true,
	"котофейки":   true,
}

type TheCatAPIDataGetter struct{}

func (*TheCatAPIDataGetter) GetData(msg string) (string, string, error) {
	imageURL, err := thecatapi.GetCatImageURL(os.Getenv("CATAPI_TOKEN"))
	if err != nil {
		log.Fatal("Can't get image from TheCatAPI", err.Error())
		return "", "", err
	}

	return "photo", imageURL, nil
}

func CatInText(text string) bool {
	for _, word := range strings.Fields(text) {
		if catVocabulary[strings.ToLower(word)] {
			return true
		}
	}

	return false
}

func HandleTelegramWebHook(w http.ResponseWriter, req *http.Request) {
	update, err := telegramapi.ParseUpdate(req)
	if err != nil {
		log.Fatal("Cant' parse update message", err.Error())
		return
	}

	userMsg := update.Message.Text

	if CatInText(userMsg) {
		chatId := update.Message.Chat.Id
		telegramResp, err := telegramapi.SendDataToTelegramChat(&TheCatAPIDataGetter{}, userMsg, chatId, telegramapi.SendPhotoMethod)

		if err != nil {
			log.Fatalf("Can't send message from CatBot to chat %d, error: %s", chatId, err.Error())
			return
		}

		log.Printf("Message was sent to chat %d, Telegram Bot API response %s", chatId, telegramResp)
	}
}

func main() {
	http.HandleFunc("/", HandleTelegramWebHook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}
