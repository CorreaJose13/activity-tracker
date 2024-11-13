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
	pinki   = "https://scontent.fclo8-1.fna.fbcdn.net/v/t39.30808-6/438734143_848452533981874_4817032528961290089_n.jpg?_nc_cat=100&ccb=1-7&_nc_sid=bd9a62&_nc_ohc=2PwwxOpwaxIQ7kNvgFy3y0R&_nc_zt=23&_nc_ht=scontent.fclo8-1.fna&_nc_gid=A7V0dhgzh2JTTbFFzNtHs3k&oh=00_AYCd9sWEqbIXqz-WtD8NvL0oSChHI8m_oBvz0FfZEofWXA&oe=673849E8"
)

var (
	goalWaterConsume    = 3
	goalKeratineConsume = 1

	commandMap = map[string]func(bot *telegram.Bot, chatID int64) error{
		"/hello":       sendHello,
		"/help":        sendHelp,
		"/commands":    sendCommands,
		"/track":       sendTrackHelp,
		"/report":      sendReportHelp,
		"/hatriki":     sendHatriki,
		"/tengohambre": sendHambre,
		"/pinkipiensa": sendPinki,
	}

	suffixReportMap = map[string]func(bot *telegram.Bot, userName, content string, chatID int64) error{
		"water":    sendWaterReport,
		"keratine": sendKeratineReport,
		"pipi":     sendPipiReport,
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
		shared.Keratine:   sendTrackKeratine,
		shared.Pipi:       sendTrackPipi,
	}
)

func doCommand(bot *telegram.Bot, chatID int64, userName string, command string) error {
	parts := strings.Split(command, " ")

	before, _, found := strings.Cut(parts[0], "@")
	if found {
		parts[0] = before
		command = strings.Join(parts, " ")
	}

	fn, ok := commandMap[command]
	if ok {
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

func isWaterGoalCompleted(bot *telegram.Bot, userName string, chatID int64) bool {
	currentDayWaterActivities, err := storage.GetCurrentDayActivities(userName, shared.Water)
	if err != nil {
		_ = telegram.SendMessage(bot, chatID, "tenemos problemas papi"+err.Error())

		return false
	}

	return len(currentDayWaterActivities) >= goalWaterConsume
}

func isKeratineGoalCompleted(bot *telegram.Bot, userName string, chatID int64) bool {
	currentDayKeratineActivities, err := storage.GetCurrentDayActivities(userName, shared.Keratine)
	if err != nil {
		_ = telegram.SendMessage(bot, chatID, "tenemos problemas papi"+err.Error())

		return false
	}

	return len(currentDayKeratineActivities) == goalKeratineConsume
}

func sendTrackWater(bot *telegram.Bot, userName, content string, chatID int64) error {
	isGoalCompleted := isWaterGoalCompleted(bot, userName, chatID)
	if isGoalCompleted {
		return telegram.SendMessage(bot, chatID, "ya te tomaste los 3L de awa mi papacho, aprende a tener límites")
	}

	now, err := shared.GetNow()
	if err != nil {
		return telegram.SendMessage(bot, chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.Water),
		Name:      userName,
		Activity:  shared.Water,
		CreatedAt: nowStr,
		Content:   content, // TODO: add logic to validate the content and use it in isGoalCompleted function
	}

	err = storage.Create(userActivity)
	if err != nil {
		return telegram.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return telegram.SendMessage(bot, chatID, "se wardó tu tomadita de awa golosito")
}

func sendTrackPipi(bot *telegram.Bot, userName, content string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return telegram.SendMessage(bot, chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.Pipi),
		Name:      userName,
		Activity:  shared.Pipi,
		CreatedAt: nowStr,
	}

	err = storage.Create(userActivity)
	if err != nil {
		return telegram.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	if userName == shared.Valery {
		return telegram.SendMessage(bot, chatID, "Epa, buena esa mionsita 😎")
	}

	return telegram.SendMessage(bot, chatID, "Epa, buena esa mionsito 😎")
}

func sendTrackKeratine(bot *telegram.Bot, userName, content string, chatID int64) error {
	isGoalCompleted := isKeratineGoalCompleted(bot, userName, chatID)
	if isGoalCompleted {
		return telegram.SendMessage(bot, chatID, "ya te tomaste la keratina de hoy, aprende a tener límites xfi")
	}

	now, err := shared.GetNow()
	if err != nil {
		return telegram.SendMessage(bot, chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.Keratine),
		Name:      userName,
		Activity:  shared.Keratine,
		CreatedAt: nowStr,
		Content:   content,
	}

	err = storage.Create(userActivity)
	if err != nil {
		return telegram.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return telegram.SendMessage(bot, chatID, "se wardó tu tomadita de keratina >:)")
}

func sendTrackTooth(bot *telegram.Bot, userName, _ string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return telegram.SendMessage(bot, chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.ToothBrush),
		Name:      userName,
		Activity:  shared.ToothBrush,
		CreatedAt: nowStr,
	}

	err = storage.Create(userActivity)
	if err != nil {
		return telegram.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return telegram.SendMessage(bot, chatID, "menos mal, ya te olia a qlo la boca mi papacho 💩")
}

func sendTrackRun(bot *telegram.Bot, userName, content string, chatID int64) error {
	now, err := shared.GetNow()
	if err != nil {
		return telegram.SendMessage(bot, chatID, err.Error())
	}

	nowStr := now.Format(time.RFC3339)

	userActivity := storage.UserActivity{
		ID:        storage.GenerateActivityItemID(now, userName, shared.ToothBrush),
		Name:      userName,
		Activity:  shared.ToothBrush,
		CreatedAt: nowStr,
		Content:   content,
	}

	err = storage.Create(userActivity)
	if err != nil {
		return telegram.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
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
	return telegram.SendMessage(bot, chatID, "higado al fallo")
}

func sendTrackPoop(bot *telegram.Bot, userName, content string, chatID int64) error {
	return telegram.SendMessage(bot, chatID, "a ber?")
}

func sendWaterReport(bot *telegram.Bot, userName, content string, chatID int64) error {
	wr, err := reports.GenerateWaterReport(bot, userName, chatID)
	if err != nil {
		return err
	}

	return telegram.SendMessage(bot, chatID, wr)
}

func sendPipiReport(bot *telegram.Bot, userName, content string, chatID int64) error {
	pr, err := reports.GeneratePipiReport(bot, userName, chatID)
	if err != nil {
		return err
	}

	return telegram.SendMessage(bot, chatID, pr)
}

func sendKeratineReport(bot *telegram.Bot, userName, content string, chatID int64) error {
	kr, err := reports.GenerateKeratineReport(bot, userName, chatID)
	if err != nil {
		return err
	}

	return telegram.SendMessage(bot, chatID, kr)
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

func sendPinki(bot *telegram.Bot, chatID int64) error {
	err := telegram.SendPhoto(bot, chatID, pinki)
	if err != nil {
		return telegram.SendMessage(bot, chatID, "vamos pinki piensa")
	}

	return nil
}
