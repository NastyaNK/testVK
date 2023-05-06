package listener

import (
	"fmt"
	"log"

	"mini_telegram_bot/keyboard"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type consumer struct {
	bot     *tgbotapi.BotAPI
	channel tgbotapi.UpdatesChannel
}

func New(token string, timeout int, debug bool) (consumer, error) {
	var listener consumer

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return listener, fmt.Errorf("create bot err: %w", err)
	}

	bot.Debug = debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	ucfg := tgbotapi.NewUpdate(0)
	ucfg.Timeout = timeout

	chanel, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		return listener, fmt.Errorf("create chanel err: %w", err)
	}

	listener.bot = bot
	listener.channel = chanel

	return listener, nil
}

func (c consumer) Listen() {
	for update := range c.channel {
		UserName := update.Message.From.UserName
		ChatID := update.Message.Chat.ID

		text := update.Message.Text
		log.Printf("[%s] %d %s", UserName, ChatID, text)

		msg := tgbotapi.NewMessage(ChatID, text)

		buttons, ok := keyboard.Create(text)
		if ok {
			if text == "/start" {
				text = "Привет " + UserName
			}

			if err := keyboard.Set(ChatID, text, buttons); err != nil {
				log.Printf("set new keyboard err: %v", err)
			}

			continue
		}

		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("sent message err: %v", err)
		}
	}
}
