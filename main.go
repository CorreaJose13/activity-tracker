package main

import (
	tgClient "activity-tracker/api/telegram"
	"activity-tracker/config"
	eventConsumer "activity-tracker/consumer/event-consumer"
	"log"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Println(err)
		return
	}
	bot, err := tgClient.New(cfg.TgBotToken)
	if err != nil {
		log.Println(err)
		return
	}
	update := tgClient.Updates(bot)
	eventConsumer.Processor(bot, update)
}
