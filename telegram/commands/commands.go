package commands

import (
	"activity-tracker/shared"
	"activity-tracker/telegram/commands/goals"
	"activity-tracker/telegram/commands/report"
	"activity-tracker/telegram/commands/track"
	"strings"
)

const (
	hatriki = "https://external-preview.redd.it/jrtz49x5F1cjvDQoFzb0I4cv2dwhA5RDhqaEcBbiXIU.png?format=pjpg&auto=webp&s=3ef741c83f7927eca91cb8ac2d610fd6f010d5b0"
	jeje    = "https://static.wikia.nocookie.net/memes-pedia/images/5/5e/Quieres_Pene.jpg/revision/latest/scale-to-width-down/1200?cb=20230507024715&path-prefix=es"
	pinki   = "https://scontent.fclo8-1.fna.fbcdn.net/v/t39.30808-6/438734143_848452533981874_4817032528961290089_n.jpg?_nc_cat=100&ccb=1-7&_nc_sid=bd9a62&_nc_ohc=2PwwxOpwaxIQ7kNvgFy3y0R&_nc_zt=23&_nc_ht=scontent.fclo8-1.fna&_nc_gid=A7V0dhgzh2JTTbFFzNtHs3k&oh=00_AYCd9sWEqbIXqz-WtD8NvL0oSChHI8m_oBvz0FfZEofWXA&oe=673849E8"
)

var (
	commandMap = map[string]func(bot *shared.Bot, chatID int64) error{
		"/hello":       sendHello,
		"/help":        sendHelp,
		"/commands":    sendCommands,
		"/track":       sendTrackHelp,
		"/report":      sendReportHelp,
		"/goal":        sendGoalHelp,
		"/hatriki":     sendHatriki,
		"/tengohambre": sendHambre,
		"/pinkipiensa": sendPinki,
	}

	suffixReportMap = map[string]func(bot *shared.Bot, userName, content string, chatID int64) error{
		"water":    report.SendWaterReport,
		"keratine": report.SendKeratineReport,
		"pipi":     report.SendPipiReport,
	}

	suffixTrackMap = map[shared.Activity]func(bot *shared.Bot, userName, content string, chatID int64) error{
		shared.Water:      track.SendTrackWater,
		shared.ToothBrush: track.SendTrackTooth,
		shared.Read:       track.SendTrackRead,
		shared.Shower:     track.SendTrackShower,
		shared.Sleep:      track.SendTrackSleep,
		shared.Gym:        sendTrackGym,
		shared.Poop:       track.SendTrackPoop,
		shared.Run:        track.SendTrackRun,
		shared.Keratine:   track.SendTrackKeratine,
		shared.Pipi:       track.SendTrackPipi,
		shared.Wishlist:   track.SendTrackWishlist,
	}

	suffixGoalMap = map[string]func(bot *telegram.Bot, userName, content string, chatID int64) error{
		"create": goals.SendCreateGoal,
		"delete": goals.SendDeleteGoal,
		"update": goals.SendUpdateGoal,
		"all":    goals.SendAllGoals,
	}

	msgHelp = `Quieres pene?`

	msgCommands = `Here's a list of commands you can send me:

	/hello
	/help
	/commands
	/track
	/report`

	msgTrack = `papi y entonces? qué te trackeo? las veces que te engañó tu ex o q, mandame info sapa. 
hint: 
-/track water
-/track toothbrush
-/track read
-/track shower
-/track sleep
-/track gym
-/track poop`

	msgHello = "Hola precioso \n\n" + msgHelp

	msgUnknownCommand = "Unknown command 🤔"
)

// DoCommand handles the command
func DoCommand(bot *shared.Bot, chatID int64, userName string, command string) error {
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

	if strings.HasPrefix(command, "/track ") {
		suffix := strings.TrimPrefix(command, "/track ")
		return handleTrack(bot, chatID, userName, suffix)
	}

	if strings.HasPrefix(command, "/report ") {
		suffix := strings.TrimPrefix(command, "/report ")
		return handleReport(bot, chatID, userName, suffix)
	}

	if strings.HasPrefix(command, "/goal ") {
		suffix := strings.TrimPrefix(command, "/goal ")
		return handleGoal(bot, chatID, userName, suffix)
	}

	return sendUnknownCommand(bot, chatID)
}

func handleTrack(bot *shared.Bot, chatID int64, userName, suffix string) error {
	before, after, _ := strings.Cut(suffix, " ")

	if fn, ok := suffixTrackMap[shared.Activity(before)]; ok {
		return fn(bot, userName, after, chatID)
	}

	return sendUnknownCommand(bot, chatID)
}

func handleReport(bot *shared.Bot, chatID int64, userName, suffix string) error {
	before, after, _ := strings.Cut(suffix, " ")

	if fn, ok := suffixReportMap[before]; ok {
		return fn(bot, userName, after, chatID)
	}

	return sendUnknownCommand(bot, chatID)
}

func handleGoal(bot *telegram.Bot, chatID int64, userName, suffix string) error {
	before, after, _ := strings.Cut(suffix, " ")

	if fn, ok := suffixGoalMap[before]; ok {
		return fn(bot, userName, after, chatID)
	}

	return sendUnknownCommand(bot, chatID)
}

func sendUnknownCommand(bot *shared.Bot, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgUnknownCommand)
}

func sendHello(bot *shared.Bot, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgHello)
}

func sendHelp(bot *shared.Bot, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgHelp)
}

func sendGoalHelp(bot *shared.Bot, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgGoal)
}

func sendCommands(bot *shared.Bot, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgCommands)
}

func sendTrackHelp(bot *shared.Bot, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgTrack)
}

func sendTrackGym(bot *shared.Bot, userName, content string, chatID int64) error {
	return shared.SendMessage(bot, chatID, "higado al fallo")
}

func sendReportHelp(bot *shared.Bot, chatID int64) error {
	return shared.SendMessage(bot, chatID, "reporthelp")
}

func sendHatriki(bot *shared.Bot, chatID int64) error {
	return shared.SendPhoto(bot, chatID, hatriki)
}

func sendHambre(bot *shared.Bot, chatID int64) error {
	return shared.SendPhoto(bot, chatID, jeje)
}

func sendPinki(bot *shared.Bot, chatID int64) error {
	err := shared.SendPhoto(bot, chatID, pinki)
	if err != nil {
		return shared.SendMessage(bot, chatID, "vamos pinki piensa")
	}

	return nil
}
