package main

import (
	"log"

	"mini_telegram_bot/listener"
)

func main() {
	consumer, err := listener.New("5907484224:AAGyf0jK4NbJwiQTQq7leBvOqChnLwV7R4Y", 60, false)
	if err != nil {
		log.Fatalf("start bot, err: %v", err)
	}

	consumer.Listen()
	log.Fatal("bot stop")
}
