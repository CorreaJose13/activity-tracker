package eventconsumer

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	tgClient "activity-tracker/api/telegram"
	"activity-tracker/events/telegram"
)

func Processor(bot *tgClient.Bot, updates tgClient.Channel) (err error) {
	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Tell the user the bot is online
	log.Println("Estamo activooooo papi, escribe cualquier mondá")

	err = telegram.Fetch(ctx, bot, updates)
	if err != nil {
		fmt.Println(err)
	}

	// Wait for a newline symbol, then cancel handling updates
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

	return err
}
