package main

import (
	"catbot/telegramapi"
	"catbot/thecatapi"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
)

var catVocabulary = map[string]bool{
	"кот":          true,
	"кота":         true,
	"коту":         true,
	"коте":         true,
	"коты":         true,
	"котэ":         true,
	"котом":        true,
	"котя":         true,
	"котю":         true,
	"коти":         true,
	"коть":         true,
	"котов":        true,
	"котейка":      true,
	"котейку":      true,
	"котейке":      true,
	"котейкой":     true,
	"котейки":      true,
	"котеек":       true,
	"котька":       true,
	"котьку":       true,
	"котьке":       true,
	"котьки":       true,
	"котькой":      true,
	"котек":        true,
	"котан":        true,
	"котана":       true,
	"котану":       true,
	"котане":       true,
	"котаны":       true,
	"котанов":      true,
	"котаном":      true,
	"котеечка":     true,
	"котеечку":     true,
	"котеечке":     true,
	"котеечки":     true,
	"котеечек":     true,
	"котеечкой":    true,
	"котейшество":  true,
	"котейшеств":   true,
	"котейшества":  true,
	"котейшеству":  true,
	"котейшестве":  true,
	"котейшеством": true,
	"котей":        true,
	"котеи":        true,
	"котея":        true,
	"котее":        true,
	"котею":        true,
	"котев":        true,
	"котеев":       true,
	"котофей":      true,
	"котофея":      true,
	"котофею":      true,
	"котофее":      true,
	"котофеи":      true,
	"котофеев":     true,
	"котофеем":     true,
	"котофейка":    true,
	"котофейку":    true,
	"котофейке":    true,
	"котофейки":    true,
	"котофеек":     true,
	"котофейкой":   true,
	"котёнок":      true,
	"котёнка":      true,
	"котёнку":      true,
	"котёнке":      true,
	"котёнки":      true,
	"котёнков":     true,
	"котёнком":     true,
	"котёночек":    true,
	"котёночка":    true,
	"котёночку":    true,
	"котёночке":    true,
	"котёночки":    true,
	"котёночков":   true,
	"котёночком":   true,
	"котик":        true,
	"котика":       true,
	"котику":       true,
	"котике":       true,
	"котики":       true,
	"котиков":      true,
	"котиком":      true,
	"котенок":      true,
	"котенка":      true,
	"котенку":      true,
	"котенке":      true,
	"котенки":      true,
	"котенков":     true,
	"котенком":     true,
	"котеночек":    true,
	"котеночка":    true,
	"котеночку":    true,
	"котеночке":    true,
	"котеночки":    true,
	"котеночков":   true,
	"котеночком":   true,
	"котяра":       true,
	"котяры":       true,
	"котяру":       true,
	"котяре":       true,
	"котярой":      true,
	"котяр":        true,
	"котища":       true,
	"котищ":        true,
	"котищей":      true,
	"котищу":       true,
	"котищи":       true,
	"котище":       true,
}

var gifVocabulary = map[string]bool{
	"гифка": true,
	"гифки": true,
	"гифку": true,
	"гифок": true,
	"гифке": true,
	"гиф":   true,
}

type TheCatAPIDataGetter struct {
	gif bool
}

func (theCatAPIDataGetter *TheCatAPIDataGetter) GetData(msg string) (string, string, error) {
	imageURL, err := thecatapi.GetCatImageURL(os.Getenv("CATAPI_TOKEN"), theCatAPIDataGetter.gif)
	if err != nil {
		log.Fatal("Can't get image from TheCatAPI", err.Error())
		return "", "", err
	}

	if theCatAPIDataGetter.gif {
		return "animation", imageURL, nil
	}

	return "photo", imageURL, nil
}

func splitText(text string) []string {
	splitFunc := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	return strings.FieldsFunc(text, splitFunc)
}

func CatInText(words []string) bool {
	for _, word := range words {
		if catVocabulary[strings.ToLower(word)] {
			return true
		}
	}

	return false
}

func GifInText(words []string) bool {
	for _, word := range words {
		if gifVocabulary[strings.ToLower(word)] {
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
	words := splitText(userMsg)

	if CatInText(words) {
		chatId := update.Message.Chat.Id
		theCatAPIDataGetter := TheCatAPIDataGetter{GifInText(words)}

		var telegramResp string
		if theCatAPIDataGetter.gif {
			telegramResp, err = telegramapi.SendDataToTelegramChat(&theCatAPIDataGetter, userMsg, chatId, telegramapi.SendAnimationMethod)
		} else {
			telegramResp, err = telegramapi.SendDataToTelegramChat(&theCatAPIDataGetter, userMsg, chatId, telegramapi.SendPhotoMethod)
		}

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
