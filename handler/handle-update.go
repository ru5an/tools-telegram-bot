package handler

import (
	"go-telegram-tools/config"
	"go-telegram-tools/model"
	"go-telegram-tools/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, upd tgbotapi.Update, cfg *config.Config, userStates map[int64]*state.UserState) {
	if upd.Message == nil {
		return
	}

	chatID := upd.Message.Chat.ID
	state := state.GetUserState(userStates, chatID)

	if upd.Message.IsCommand() {
		HandleCommand(bot, chatID, upd.Message.Command())
		return
	}

	switch state.Mode {
	case model.MainMode:
		HandleMainMode(bot, state, chatID, upd.Message.Text)
	case model.PdfMode:
		HandlePdfMode(bot, state, chatID, upd, cfg)
	default:
		SendMainMenu(bot, chatID)
	}
}
