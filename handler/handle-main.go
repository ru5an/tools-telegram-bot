package handler

import (
	"go-telegram-tools/model"
	"go-telegram-tools/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMainMode(bot *tgbotapi.BotAPI, state *state.UserState, chatID int64, text string) {
	switch text {
	case "Робота з PDF":
		state.Mode = model.PdfMode
		SendPdfMenu(bot, chatID)
	default:
		SendMainMenu(bot, chatID)
	}
}

func SendMainMenu(bot *tgbotapi.BotAPI, chatID int64) {
	model.SendWithKeyboard(bot, chatID, model.MainKeyboard, "Оберіть дію")
}
