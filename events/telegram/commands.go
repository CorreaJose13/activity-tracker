package telegram

import (
	"activity-tracker/api/telegram"
)

const (
	hello    = "/hello"
	help     = "/help"
	commands = "/commands"
	track    = "/trackprogress"
	water    = "water"
	tooth    = "toothbrush"
	read     = "reading"
	shower   = "shower"
	sleep    = "sleep"
	gym      = "gym"
	poop     = "shit"
	rep      = "report"
	all      = "all"
)

func doCommand(bot *telegram.Bot, chatID int64, command string) (err error) {

	switch command {
	case hello:
		return sendHello(bot, chatID)
	case help:
		return sendHelp(bot, chatID)
	case commands:
		return sendCommands(bot, chatID)
	case track:
		return sendTrackHelp(bot, chatID)
	//line 35-48 will be replaced with a function that handles /trackprogress + suffix
	case track + " " + water:
		return telegram.SendMessage(bot, chatID, "awa")
	case track + " " + tooth:
		return telegram.SendMessage(bot, chatID, "tooth")
	case track + " " + read:
		return telegram.SendMessage(bot, chatID, "read")
	case track + " " + shower:
		return telegram.SendMessage(bot, chatID, "Ya era hora")
	case track + " " + sleep:
		return telegram.SendMessage(bot, chatID, "A mimir")
	case track + " " + gym:
		return telegram.SendMessage(bot, chatID, "Higado al fallo")
	case track + " " + poop:
		return telegram.SendMessage(bot, chatID, "y la foto?")
	case "/hatriki":
		return telegram.SendPhoto(bot, chatID)
	default:
		return telegram.SendMessage(bot, chatID, msgUnknownCommand)
	}

}

func sendHello(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, msgHello)
}

func sendHelp(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, msgHelp)
}

func sendCommands(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, msgCommands)
}

func sendTrackHelp(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, msgTrack)
}
