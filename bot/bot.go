package bot

import (
	"go-telegram-tools/config"
	"go-telegram-tools/handler"
	"go-telegram-tools/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(cfg *config.Config, userStates map[int64]*state.UserState) error {
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	upds := bot.GetUpdatesChan(u)

	for upd := range upds {
		handler.HandleUpdate(bot, upd, cfg, userStates)
	}

	return nil
}
