package telegram

import (
	"activity-tracker/api/telegram"
	"activity-tracker/reports"
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"strings"
	"time"
)

const (
	hatriki = "https://external-preview.redd.it/jrtz49x5F1cjvDQoFzb0I4cv2dwhA5RDhqaEcBbiXIU.png?format=pjpg&auto=webp&s=3ef741c83f7927eca91cb8ac2d610fd6f010d5b0"
	jeje    = "https://static.wikia.nocookie.net/memes-pedia/images/5/5e/Quieres_Pene.jpg/revision/latest/scale-to-width-down/1200?cb=20230507024715&path-prefix=es"
)

var (
	goalWaterConsume = 3

	commandMap = map[string]func(bot *telegram.Bot, chatID int64) error{
		"/hello":       sendHello,
		"/help":        sendHelp,
		"/commands":    sendCommands,
		"/track":       sendTrackHelp,
		"/report":      sendReportHelp,
		"/hatriki":     sendHatriki,
		"/tengohambre": sendHambre,
	}

	suffixReportMap = map[string]func(bot *telegram.Bot, userName string, chatID int64) error{
		"water": sendWaterReport,
		// add the other report commands here when they are implemented
	}

	suffixTrackMap = map[string]func(bot *telegram.Bot, userName string, chatID int64) error{
		"water":      sendTrackWater,
		"toothbrush": sendTrackTooth,
		"read":       sendTrackRead,
		"shower":     sendTrackShower,
		"sleep":      sendTrackSleep,
		"gym":        sendTrackGym,
		"poop":       sendTrackPoop,
	}
)

func doCommand(bot *telegram.Bot, chatID int64, userName string, command string) error {
	if fn, ok := commandMap[command]; ok {
		return fn(bot, chatID)
	}

	// Check if the command starts with "track"
	if strings.HasPrefix(command, "/track ") {
		suffix := strings.TrimPrefix(command, "/track ")
		return handleTrack(bot, chatID, userName, suffix)
	}

	if strings.HasPrefix(command, "/report ") {
		suffix := strings.TrimPrefix(command, "/report ")
		return handleReport(bot, chatID, userName, suffix)
	}

	return sendUnknownCommand(bot, chatID)
}

func handleTrack(bot *telegram.Bot, chatID int64, userName, suffix string) error {
	if fn, ok := suffixTrackMap[suffix]; ok {
		return fn(bot, userName, chatID)
	}

	return sendUnknownCommand(bot, chatID)
}

func handleReport(bot *telegram.Bot, chatID int64, userName, suffix string) error {
	if fn, ok := suffixReportMap[suffix]; ok {
		return fn(bot, userName, chatID)
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

func isGoalCompleted(bot *telegram.Bot, userName string, chatID int64) bool {
	currentDayWaterActivities, err := storage.GetCurrentDayActivities(userName, shared.Water)
	if err != nil {
		_ = telegram.SendMessage(bot, chatID, "tenemos problemas papi"+err.Error())

		return false
	}

	return len(currentDayWaterActivities) >= goalWaterConsume
}

func sendTrackWater(bot *telegram.Bot, userName string, chatID int64) error {
	isGoalCompleted := isGoalCompleted(bot, userName, chatID)
	if isGoalCompleted {
		return telegram.SendMessage(bot, chatID, "ya te tomaste los 3L de awa mi papacho, aprende a tener límites")
	}

	now := time.Now()

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.Water),
		Name:      userName,
		Activity:  shared.Water,
		CreatedAt: now,
	}

	err := storage.Create(userActivity)
	if err != nil {
		return telegram.SendMessage(bot, chatID, "algo falló mi fafá: "+err.Error())
	}

	return telegram.SendMessage(bot, chatID, "se wardó tu tomadita de awa golosito")
}

func sendTrackTooth(bot *telegram.Bot, userName string, chatID int64) error {
	now := time.Now()

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.ToothBrush),
		Name:      userName,
		Activity:  shared.ToothBrush,
		CreatedAt: now,
	}

	err := storage.Create(userActivity)
	if err != nil {
		return telegram.SendMessage(bot, chatID, "algo falló mi fafá: "+err.Error())
	}

	return telegram.SendMessage(bot, chatID, "menos mal, ya te olia a qlo la boca mi papacho 💩")
}

func sendTrackRead(bot *telegram.Bot, userName string, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "lectura")
}

func sendTrackShower(bot *telegram.Bot, userName string, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "ya era hora olias a obo")
}

func sendTrackSleep(bot *telegram.Bot, userName string, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "zzzzz")
}

func sendTrackGym(bot *telegram.Bot, userName string, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "higado al fallo")
}

func sendTrackPoop(bot *telegram.Bot, userName string, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "a ber?")
}

func sendWaterReport(bot *telegram.Bot, userName string, chatID int64) error {
	wr, err := reports.GenerateWaterReport(bot, userName, chatID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return telegram.SendMessage(bot, chatID, wr)
}

func sendReportHelp(bot *telegram.Bot, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "reporthelp")
}

func sendHatriki(bot *telegram.Bot, chatID int64) error {
	return telegram.SendPhoto(bot, chatID, hatriki)
}

func sendHambre(bot *telegram.Bot, chatID int64) error {
	return telegram.SendPhoto(bot, chatID, jeje)
}
