package utils

import (
	"io"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DownloadFile(bot *tgbotapi.BotAPI, fileID string) ([]byte, error) {
	file, _ := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	url := file.Link(bot.Token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
