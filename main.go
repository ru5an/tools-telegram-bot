package main

import (
	"go-telegram-tools/bot"
	"go-telegram-tools/config"
	"go-telegram-tools/state"
	"log"
)

// type ChatMode int

// const (
// 	MainMode ChatMode = iota
// 	PdfMode
// )

// type UserState struct {
// 	Mode  ChatMode
// 	Files [][]byte
// }

// var (
// 	conf       *config.Config
// 	userStates = make(map[int64]*UserState)
// 	mu         sync.Mutex

// 	mainKeyboard = tgbotapi.NewReplyKeyboard(
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("Робота з PDF"),
// 		),
// 	)

// 	pdfKeyboard = tgbotapi.NewReplyKeyboard(
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("Об'єднати"),
// 			tgbotapi.NewKeyboardButton("Очистити"),
// 		),
// 		tgbotapi.NewKeyboardButtonRow(
// 			tgbotapi.NewKeyboardButton("Назад"),
// 		),
// 	)
// )

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("not have env: ", err)
	}

	userStates := make(map[int64]*state.UserState)

	if err := bot.Start(conf, userStates); err != nil {
		log.Fatal("bot error: ", err)
	}

	// bot, err := tgbotapi.NewBotAPI(conf.BotToken)
	// if err != nil {
	// 	log.Fatal(err, conf.BotToken, "!")
	// }

	// updates := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))
	// handleUpdates(bot, updates)
}

// func getUserState(chatID int64) *UserState {
// 	state, exists := userStates[chatID]
// 	if !exists {
// 		state = &UserState{Mode: MainMode, Files: [][]byte{}}
// 		userStates[chatID] = state
// 	}
// 	return state
// }

// func handleUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}
// 		processMessage(bot, update)
// 	}
// }

// func processMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
// 	chatID := update.Message.Chat.ID
// 	state := getUserState(chatID)

// 	if update.Message.IsCommand() {
// 		handleCommand(bot, chatID, update.Message.Command())
// 		return
// 	}

// 	switch state.Mode {
// 	case MainMode:
// 		handleMainMode(bot, chatID, update.Message.Text)
// 	case PdfMode:
// 		handlePdfMode(bot, chatID, update)
// 	default:
// 		sendMainMenu(bot, chatID)
// 	}
// }

// func handleMainMode(bot *tgbotapi.BotAPI, chatID int64, text string) {
// 	switch text {
// 	case "Робота з PDF":
// 		state := getUserState(chatID)
// 		state.Mode = PdfMode
// 		sendPdfMenu(bot, chatID)
// 	default:
// 		sendMainMenu(bot, chatID)
// 	}
// }

// func sendMainMenu(bot *tgbotapi.BotAPI, chatID int64) {
// 	sendWithKeyboard(bot, chatID, &mainKeyboard, "Оберіть дію")
// }

// func handlePdfMode(bot *tgbotapi.BotAPI, chatID int64, update tgbotapi.Update) {
// 	state := getUserState(chatID)

// 	if update.Message.Text == "Назад" {
// 		state.Mode = MainMode
// 		sendMainMenu(bot, chatID)
// 		return
// 	}

// 	if update.Message.Text == "Очистити" {
// 		mu.Lock()
// 		state.Files = nil
// 		mu.Unlock()
// 		sendWithKeyboard(bot, chatID, &pdfKeyboard, "Файли очищено")
// 		return
// 	}

// 	if update.Message.Text == "Об'єднати" {
// 		if len(state.Files) < 2 {
// 			sendWithKeyboard(bot, chatID, &pdfKeyboard, "Недостатньо файлів")
// 			return
// 		}

// 		merged, err := pdfmanager.SendToMergeService(state.Files, conf.PdfPath+conf.MergePdfPath)

// 		if err != nil {
// 			sendWithKeyboard(bot, chatID, &pdfKeyboard, "Помилка при об'єднанні")
// 			return
// 		}

// 		doc := tgbotapi.NewDocument(chatID, tgbotapi.FileBytes{Name: "merged.pdf", Bytes: merged})
// 		sendWithKeyboard(bot, chatID, &pdfKeyboard, "Об'єднаний файл:")
// 		bot.Send(doc)

// 		state.Files = nil
// 		return
// 	}

// 	if update.Message.Document != nil && update.Message.Document.MimeType == "application/pdf" {
// 		data, err := downloadFile(bot, update.Message.Document.FileID)
// 		if err != nil {
// 			sendWithKeyboard(bot, chatID, &pdfKeyboard, "Не вдалося завантажити файл")
// 		}
// 		mu.Lock()
// 		state.Files = append(state.Files, data)
// 		mu.Unlock()
// 		sendWithKeyboard(bot, chatID, &pdfKeyboard, fmt.Sprintf("Файл додано. Всього %d файлів", len(state.Files)))
// 		return
// 	}
// }

// func sendPdfMenu(bot *tgbotapi.BotAPI, chatID int64) {
// 	sendWithKeyboard(bot, chatID, &pdfKeyboard, "Режим роботи з PDF файлами. Завантажте файли та оберіть дію.")
// }

// func handleCommand(bot *tgbotapi.BotAPI, chatID int64, cmd string) {
// 	switch cmd {
// 	case "start":
// 		sendMainMenu(bot, chatID)
// 	default:
// 		sendWithKeyboard(bot, chatID, &mainKeyboard, "Незрозуміла команда")
// 	}
// }

// func downloadFile(bot *tgbotapi.BotAPI, fileID string) ([]byte, error) {
// 	file, _ := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
// 	url := file.Link(bot.Token)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	return io.ReadAll(resp.Body)
// }

// // func sendToMergeService(pdfs [][]byte) ([]byte, error) {
// // 	var buf bytes.Buffer
// // 	writer := multipart.NewWriter(&buf)

// // 	for i, file := range pdfs {
// // 		fw, _ := writer.CreateFormFile("files", fmt.Sprintf("file%d.pdf", i+1))
// // 		fw.Write(file)
// // 	}

// // 	writer.Close()

// // 	req, _ := http.NewRequest("POST", conf.PdfPath+conf.MergePdfPath, &buf)
// // 	req.Header.Set("Content-Type", writer.FormDataContentType())

// // 	resp, err := http.DefaultClient.Do(req)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	defer resp.Body.Close()

// // 	return io.ReadAll(resp.Body)
// // }

// func sendWithKeyboard(bot *tgbotapi.BotAPI, chatID int64, keyboard *tgbotapi.ReplyKeyboardMarkup, text string) {
// 	msg := tgbotapi.NewMessage(chatID, text)
// 	msg.ReplyMarkup = keyboard
// 	bot.Send(msg)
// }
