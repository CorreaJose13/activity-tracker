package event_consumer

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"activity-tracker/api/telegram"
	event "activity-tracker/events/telegram"
)

func Processor(bot *telegram.Bot, updates telegram.Channel) error {
	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Tell the user the bot is online
	log.Println("Estamo activooooo papi, escribe cualquier mondá")

	err := event.Fetch(ctx, bot, updates)
	if err != nil {
		//a este punto pueden llegar 2 tipos de errores según el trace que llevo: error que arroja el bot.Send
		//del tgbotapi, y errMissingUser
		fmt.Println(err)
	}

	// Wait for a newline symbol, then cancel handling updates
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

	return err
}
