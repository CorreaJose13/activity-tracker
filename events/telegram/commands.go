package telegram

import (
	"activity-tracker/api/telegram"
	"activity-tracker/reports"
	"activity-tracker/shared"
	"activity-tracker/storage"

	flag "github.com/spf13/pflag"

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

	suffixReportMap = map[string]func(bot *telegram.Bot, userName, content string, chatID int64) error{
		"water": sendWaterReport,
		// add the other report commands here when they are implemented
	}

	suffixTrackMap = map[shared.Activity]func(bot *telegram.Bot, userName, content string, chatID int64) error{
		shared.Water:      sendTrackWater,
		shared.ToothBrush: sendTrackTooth,
		shared.Read:       sendTrackRead,
		shared.Shower:     sendTrackShower,
		shared.Sleep:      sendTrackSleep,
		shared.Gym:        sendTrackGym,
		shared.Poop:       sendTrackPoop,
		shared.Run:        sendTrackRun,
	}
)

func doCommand(bot *telegram.Bot, chatID int64, userName string, command string) error {
	parts := strings.Split(command, " ")

	before, _, found := strings.Cut(parts[0], "@")
	if found {
		parts[0] = before
		command = strings.Join(parts, " ")
	}

	if fn, ok := commandMap[command]; ok {
		return fn(bot, chatID)
	}

	// Check if the command starts with "track"
	if strings.HasPrefix(command, "/track ") {
		suffix := strings.TrimPrefix(command, "/track ")
		return handleTrack(bot, chatID, userName, suffix)
	}

	// Check if the command starts with "report"
	if strings.HasPrefix(command, "/report ") {
		suffix := strings.TrimPrefix(command, "/report ")
		return handleReport(bot, chatID, userName, suffix)
	}

	return sendUnknownCommand(bot, chatID)
}

func handleTrack(bot *telegram.Bot, chatID int64, userName, suffix string) error {
	before, after, _ := strings.Cut(suffix, " ")

	if fn, ok := suffixTrackMap[shared.Activity(before)]; ok {
		return fn(bot, userName, after, chatID)
	}

	return sendUnknownCommand(bot, chatID)
}

func handleReport(bot *telegram.Bot, chatID int64, userName, suffix string) error {
	before, after, _ := strings.Cut(suffix, " ")

	if fn, ok := suffixReportMap[before]; ok {
		return fn(bot, userName, after, chatID)
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

func sendTrackWater(bot *telegram.Bot, userName, content string, chatID int64) error {
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
		Content:   content, // TODO: add logic to validate the content and use it in isGoalCompleted function
	}

	err := storage.Create(userActivity)
	if err != nil {
		return telegram.SendMessage(bot, chatID, "algo falló mi fafá: "+err.Error())
	}

	return telegram.SendMessage(bot, chatID, "se wardó tu tomadita de awa golosito")
}

func sendTrackTooth(bot *telegram.Bot, userName, _ string, chatID int64) error {
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

func sendTrackRun(bot *telegram.Bot, userName, content string, chatID int64) error {
	now := time.Now()

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.ToothBrush),
		Name:      userName,
		Activity:  shared.ToothBrush,
		CreatedAt: now,
		Content:   content,
	}

	err := storage.Create(userActivity)
	if err != nil {
		return telegram.SendMessage(bot, chatID, "algo falló mi fafá: "+err.Error())
	}

	message := "mi papacho el más usain vol 🏃‍♂️"
	if content != "" {
		message = fmt.Sprintf("uy mi papacho corrió %s? lo iba robar un negro o qué manito. anwy congrats", content)
	}

	return telegram.SendMessage(bot, chatID, message)
}

func sendTrackRead(bot *telegram.Bot, userName, content string, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "lectura")
}

func sendTrackShower(bot *telegram.Bot, userName, content string, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "ya era hora olias a obo")
}

func sendTrackSleep(bot *telegram.Bot, userName, content string, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "zzzzz")
}

func sendTrackGym(bot *telegram.Bot, userName, content string, chatID int64) error {
	now := time.Now()

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.Gym),
		Name:      userName,
		Activity:  shared.Gym,
		CreatedAt: now,
		Content:   content,
	}

	var time string
	var muscle string

	flag.StringVarP(&time, "time", "t", "", "time you were exercising")
	flag.StringVarP(&muscle, "muscle", "p", "", "muscles you exercised splitted by comma [bicep,back,shoulder]")
	flag.Parse()

	if time == "" {
		return telegram.SendMessage(bot, chatID, "eh pero cuánto tiempo te ejercitaste sapa inmunda, mandame el -time")
	}

	if muscle == "" {
		return telegram.SendMessage(bot, chatID, "eh pero hiciste chisme al fallo o q 🐸? mandame el -muscle sapa. Ej: -muscle bicep,pecho,jeta")
	}

	err := storage.Create(userActivity)
	if err != nil {
		return telegram.SendMessage(bot, chatID, "algo falló mi fafá: "+err.Error())
	}

	return telegram.SendMessage(bot, chatID, "isss mi papacho el próximo cbum ve")
}

func sendTrackPoop(bot *telegram.Bot, userName, content string, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "a ber?")
}

func sendWaterReport(bot *telegram.Bot, userName, content string, chatID int64) error {
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
