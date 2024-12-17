package commands

import (
	"activity-tracker/shared"
	"activity-tracker/telegram/commands/gemini"
	"activity-tracker/telegram/commands/goals"
	"activity-tracker/telegram/commands/report"
	"activity-tracker/telegram/commands/track"
	"activity-tracker/telegram/commands/wishlist"
	"fmt"
	"strings"
)

const (
	hatriki = "https://external-preview.redd.it/jrtz49x5F1cjvDQoFzb0I4cv2dwhA5RDhqaEcBbiXIU.png?format=pjpg&auto=webp&s=3ef741c83f7927eca91cb8ac2d610fd6f010d5b0"
	jeje    = "https://static.wikia.nocookie.net/memes-pedia/images/5/5e/Quieres_Pene.jpg/revision/latest/scale-to-width-down/1200?cb=20230507024715&path-prefix=es"
	pinki   = "https://scontent.fclo8-1.fna.fbcdn.net/v/t39.30808-6/438734143_848452533981874_4817032528961290089_n.jpg?_nc_cat=100&ccb=1-7&_nc_sid=bd9a62&_nc_ohc=2PwwxOpwaxIQ7kNvgFy3y0R&_nc_zt=23&_nc_ht=scontent.fclo8-1.fna&_nc_gid=A7V0dhgzh2JTTbFFzNtHs3k&oh=00_AYCd9sWEqbIXqz-WtD8NvL0oSChHI8m_oBvz0FfZEofWXA&oe=673849E8"
)

var (
	commandMap = map[string]func(bot *shared.Bot, userName string, chatID int64) error{
		"/hello":       sendHello,
		"/help":        sendHelp,
		"/commands":    sendCommands,
		"/track":       sendTrackHelp,
		"/report":      sendReportHelp,
		"/goal":        sendGoalHelp,
		"/wishlist":    sendWishlist,
		"/hatriki":     sendHatriki,
		"/tengohambre": sendHambre,
		"/pinkipiensa": sendPinki,
		"/chatID":      sendChatID,
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
		shared.Gym:        track.SendTrackGym,
		shared.Poop:       track.SendTrackPoop,
		shared.Run:        track.SendTrackRun,
		shared.Keratine:   track.SendTrackKeratine,
		shared.Pipi:       track.SendTrackPipi,
		shared.Swimming:   track.SendTrackSwimming,
	}

	suffixGoalMap = map[string]func(bot *shared.Bot, userName, content string, chatID int64) error{
		"create": goals.SendCreateGoal,
		"delete": goals.SendDeleteGoal,
		"update": goals.SendUpdateGoal,
		"all":    goals.SendAllGoals,
	}

	prefixHandlers = map[string]func(bot *shared.Bot, chatID int64, userName, content string) error{
		"/track":    handleTrack,
		"/report":   handleReport,
		"/goal":     handleGoal,
		"/wishlist": wishlist.HandleWishlist,
		"/gemini":   gemini.HandleGemini,
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
		return fn(bot, userName, chatID)
	}

	for prefix, handler := range prefixHandlers {
		if strings.HasPrefix(command, prefix+" ") {
			content := strings.TrimPrefix(command, prefix+" ")

			return handler(bot, chatID, userName, content)
		}
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

func handleGoal(bot *shared.Bot, chatID int64, userName, suffix string) error {
	before, after, _ := strings.Cut(suffix, " ")

	if fn, ok := suffixGoalMap[before]; ok {
		return fn(bot, userName, after, chatID)
	}

	return sendUnknownCommand(bot, chatID)
}

func sendUnknownCommand(bot *shared.Bot, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgUnknownCommand)
}

func sendHello(bot *shared.Bot, userName string, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgHello)
}

func sendHelp(bot *shared.Bot, userName string, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgHelp)
}

func sendGoalHelp(bot *shared.Bot, userName string, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgGoal)
}

func sendWishlist(bot *shared.Bot, userName string, chatID int64) error {
	return wishlist.GetWishlist(bot, userName, chatID)
}

func sendCommands(bot *shared.Bot, userName string, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgCommands)
}

func sendTrackHelp(bot *shared.Bot, userName string, chatID int64) error {
	return shared.SendMessage(bot, chatID, msgTrack)
}

func sendReportHelp(bot *shared.Bot, username string, chatID int64) error {
	return shared.SendMessage(bot, chatID, "reporthelp")
}

func sendHatriki(bot *shared.Bot, userName string, chatID int64) error {
	return shared.SendPhoto(bot, chatID, hatriki)
}

func sendHambre(bot *shared.Bot, userName string, chatID int64) error {
	return shared.SendPhoto(bot, chatID, jeje)
}

func sendPinki(bot *shared.Bot, userName string, chatID int64) error {
	err := shared.SendPhoto(bot, chatID, pinki)
	if err != nil {
		return shared.SendMessage(bot, chatID, "vamos pinki piensa")
	}

	return nil
}

func sendChatID(bot *shared.Bot, userName string, chatID int64) error {
	message := fmt.Sprintf("Chat ID: %d", chatID)
	return shared.SendMessage(bot, chatID, message)
}
