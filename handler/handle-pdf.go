package handler

import (
	"fmt"
	"go-telegram-tools/config"
	"go-telegram-tools/model"
	pdfmanager "go-telegram-tools/pdf-manager"
	"go-telegram-tools/state"
	"go-telegram-tools/utils"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var mu sync.Mutex

func HandlePdfMode(bot *tgbotapi.BotAPI, state *state.UserState, chatID int64, update tgbotapi.Update, conf *config.Config) {

	if update.Message.Text == "Назад" {
		state.Mode = model.MainMode
		SendMainMenu(bot, chatID)
		return
	}

	if update.Message.Text == "Очистити" {
		mu.Lock()
		state.Files = nil
		mu.Unlock()
		model.SendWithKeyboard(bot, chatID, model.PdfKeyboard, "Файли очищено")
		return
	}

	if update.Message.Text == "Об'єднати" {
		if len(state.Files) < 2 {
			model.SendWithKeyboard(bot, chatID, model.PdfKeyboard, "Недостатньо файлів")
			return
		}

		merged, err := pdfmanager.SendToMergeService(state.Files, conf.PdfPath+conf.MergePdfPath)

		if err != nil {
			model.SendWithKeyboard(bot, chatID, model.PdfKeyboard, "Помилка при об'єднанні")
			return
		}

		doc := tgbotapi.NewDocument(chatID, tgbotapi.FileBytes{Name: "merged.pdf", Bytes: merged})
		model.SendWithKeyboard(bot, chatID, model.PdfKeyboard, "Об'єднаний файл:")
		bot.Send(doc)

		state.Files = nil
		return
	}

	if update.Message.Document != nil && update.Message.Document.MimeType == "application/pdf" {
		data, err := utils.DownloadFile(bot, update.Message.Document.FileID)
		if err != nil {
			model.SendWithKeyboard(bot, chatID, model.PdfKeyboard, "Не вдалося завантажити файл")
		}
		mu.Lock()
		state.Files = append(state.Files, data)
		mu.Unlock()
		model.SendWithKeyboard(bot, chatID, model.PdfKeyboard, fmt.Sprintf("Файл додано. Всього %d файлів", len(state.Files)))
		return
	}
}

func SendPdfMenu(bot *tgbotapi.BotAPI, chatID int64) {
	model.SendWithKeyboard(bot, chatID, model.PdfKeyboard, "Режим роботи з PDF файлами. Завантажте файли та оберіть дію.")
}
