package keyboard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const baseURL = "https://api.telegram.org/bot5907484224:AAGyf0jK4NbJwiQTQq7leBvOqChnLwV7R4Y/"

type sendMessageRequest struct {
	ChatID      int64                        `json:"chat_id"`
	Text        string                       `json:"text"`
	ReplyMarkup tgbotapi.ReplyKeyboardMarkup `json:"reply_markup"`
}

func Set(chatID int64, text string, buttons tgbotapi.ReplyKeyboardMarkup) error {
	request, err := json.Marshal(sendMessageRequest{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: buttons,
	})
	if err != nil {
		return fmt.Errorf("marshal err: %w", err)
	}

	response, err := http.Post(baseURL+"sendMessage", "application/json", bytes.NewBuffer(request))
	if err != nil {
		return fmt.Errorf("send err: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed")
	}

	return nil
}

func Create(msg string) (tgbotapi.ReplyKeyboardMarkup, bool) {
	var buttons tgbotapi.ReplyKeyboardMarkup
	ok := true

	switch {
	case msg == "/start" || msg == "назад" || msg == "наверх":
		buttons = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Ввести номер"),
				tgbotapi.NewKeyboardButton("Ввести имя"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Ввести данные"),
				tgbotapi.NewKeyboardButton("Ввести дату"),
			),
		)
	case msg == "Ввести номер" || msg == "Ввести имя" || msg == "Ввести данные" || msg == "Ввести дату":
		buttons = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("назад"),
				tgbotapi.NewKeyboardButton("наверх"),
			),
		)
	default:
		ok = false
	}

	return buttons, ok
}
