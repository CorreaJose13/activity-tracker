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
	hatriki  = "https://external-preview.redd.it/jrtz49x5F1cjvDQoFzb0I4cv2dwhA5RDhqaEcBbiXIU.png?format=pjpg&auto=webp&s=3ef741c83f7927eca91cb8ac2d610fd6f010d5b0"
	jeje     = "https://static.wikia.nocookie.net/memes-pedia/images/5/5e/Quieres_Pene.jpg/revision/latest/scale-to-width-down/1200?cb=20230507024715&path-prefix=es"
)

func doCommand(bot *telegram.Bot, chatID int64, command string) error {

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
		return telegram.SendPhoto(bot, chatID, hatriki)
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
