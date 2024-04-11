package telegram

import (
	"activity-tracker/api/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func doCommand(bot *tgbotapi.BotAPI, chatId int64, command string) (err error) {

	switch command {
	case "/holi":
		err = telegram.SendMessage(bot, chatId, "q dicen los hijueputaaaaaaaas")
	case "/uwu":
		err = telegram.SendMessage(bot, chatId, "ee")
	default:
		err = telegram.SendMessage(bot, chatId, "aa")
	}

	return err
}
