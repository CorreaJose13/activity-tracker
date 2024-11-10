package event_consumer

import (
	"fmt"
	"log"

	"activity-tracker/api/telegram"
	event "activity-tracker/events/telegram"
)

func Processor(bot *telegram.Bot, update telegram.Update) error {
	log.Println("Bot is running")

	err := event.Fetch(bot, update)
	if err != nil {
		//a este punto pueden llegar 2 tipos de errores según el trace que llevo: error que arroja el bot.Send
		//del tgbotapi, y errMissingUser
		fmt.Println(err)
		return err
	}

	return nil
}
