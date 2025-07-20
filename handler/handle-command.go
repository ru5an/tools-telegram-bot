package handler

import (
	"go-telegram-tools/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(bot *tgbotapi.BotAPI, chatID int64, cmd string) {
	switch cmd {
	case "start":
		SendMainMenu(bot, chatID)
	default:
		model.SendWithKeyboard(bot, chatID, model.MainKeyboard, "Незрозуміла команда")
	}
}
