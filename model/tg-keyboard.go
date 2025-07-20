package model

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	MainKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Робота з PDF"),
		),
	)

	PdfKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Об'єднати"),
			tgbotapi.NewKeyboardButton("Очистити"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
)

func SendWithKeyboard(bot *tgbotapi.BotAPI, chatID int64, keyboard tgbotapi.ReplyKeyboardMarkup, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}
