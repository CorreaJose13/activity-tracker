package telegram

import (
	"activity-tracker/api/telegram"
	"strings"
)

const (
	hatriki = "https://external-preview.redd.it/jrtz49x5F1cjvDQoFzb0I4cv2dwhA5RDhqaEcBbiXIU.png?format=pjpg&auto=webp&s=3ef741c83f7927eca91cb8ac2d610fd6f010d5b0"
	jeje    = "https://static.wikia.nocookie.net/memes-pedia/images/5/5e/Quieres_Pene.jpg/revision/latest/scale-to-width-down/1200?cb=20230507024715&path-prefix=es"
)

var commandMap = map[string]func(bot *telegram.Bot, chatID int64) error{
	"/hello":       sendHello,
	"/help":        sendHelp,
	"/commands":    sendCommands,
	"/track":       sendTrackHelp,
	"/report":      sendReportHelp,
	"/hatriki":     sendHatriki,
	"/tengohambre": sendHambre,
}

var suffixTrackMap = map[string]func(bot *telegram.Bot, chatID int64) error{
	"water":      sendTrackWater,
	"toothbrush": sendTrackTooth,
	"read":       sendTrackRead,
	"shower":     sendTrackShower,
	"sleep":      sendTrackSleep,
	"gym":        sendTrackGym,
	"poop":       sendTrackPoop,
}

func doCommand(bot *telegram.Bot, chatID int64, command string) error {
	if fn, ok := commandMap[command]; ok {
		return fn(bot, chatID)
	}
	// Check if the command starts with "track"
	if strings.HasPrefix(command, "/track ") {
		suffix := strings.TrimPrefix(command, "/track ")
		return handleTrack(bot, chatID, suffix)
	}
	return sendUnknownCommand(bot, chatID)
}

func handleTrack(bot *telegram.Bot, chatID int64, suffix string) error {
	if fn, ok := suffixTrackMap[suffix]; ok {
		return fn(bot, chatID)
	}
	return sendUnknownCommand(bot, chatID)
}

func sendUnknownCommand(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, msgUnknownCommand)
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

func sendTrackWater(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "awa")
}

func sendTrackTooth(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "diente")
}

func sendTrackRead(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "lectura")
}

func sendTrackShower(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "ya era hora olias a obo")
}

func sendTrackSleep(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "zzzzz")
}

func sendTrackGym(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "higado al fallo")
}

func sendTrackPoop(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "a ber?")
}

func sendReportHelp(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "reporthelp")
}

// modificar argumentos
func sendReportTask(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "")
}

// modificar argumentos
func sendReportAll(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "")
}

//eliminar luego

func sendHatriki(bot *telegram.Bot, chatID int64) error {
	return telegram.SendPhoto(bot, chatID, hatriki)
}

func sendHambre(bot *telegram.Bot, chatID int64) error {
	return telegram.SendPhoto(bot, chatID, jeje)
}
