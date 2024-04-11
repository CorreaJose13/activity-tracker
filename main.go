package main

import (
	"log"

	tgClient "activity-tracker/api/telegram"
	"activity-tracker/config"
	eventConsumer "activity-tracker/consumer/event-consumer"
)

func main() {

	cfg := config.MustLoad()

	bot, err := tgClient.New(cfg.TgBotToken)
	if err != nil {
		log.Println(err)
	}

	update := tgClient.Updates(bot)

	if err := eventConsumer.Processor(bot, update); err != nil {
		log.Println(err)
	}

}
