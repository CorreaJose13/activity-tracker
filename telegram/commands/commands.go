package commands

import (
	"activity-tracker/shared"
	"activity-tracker/telegram/commands/gemini"
	goals "activity-tracker/telegram/commands/goals"
	"activity-tracker/telegram/commands/report"
	shoulddeploy "activity-tracker/telegram/commands/shoulddeploy"
	"activity-tracker/telegram/commands/track"

	"activity-tracker/telegram/commands/wishlist"
	"context"
	"fmt"
	"strings"
)

const (
	hatriki = "https://external-preview.redd.it/jrtz49x5F1cjvDQoFzb0I4cv2dwhA5RDhqaEcBbiXIU.png?format=pjpg&auto=webp&s=3ef741c83f7927eca91cb8ac2d610fd6f010d5b0"
	jeje    = "https://static.wikia.nocookie.net/memes-pedia/images/5/5e/Quieres_Pene.jpg/revision/latest/scale-to-width-down/1200?cb=20230507024715&path-prefix=es"
	pinki   = "https://scontent.fclo8-1.fna.fbcdn.net/v/t39.30808-6/438734143_848452533981874_4817032528961290089_n.jpg?_nc_cat=100&ccb=1-7&_nc_sid=bd9a62&_nc_ohc=2PwwxOpwaxIQ7kNvgFy3y0R&_nc_zt=23&_nc_ht=scontent.fclo8-1.fna&_nc_gid=A7V0dhgzh2JTTbFFzNtHs3k&oh=00_AYCd9sWEqbIXqz-WtD8NvL0oSChHI8m_oBvz0FfZEofWXA&oe=673849E8"
)

var (
	commandMap = map[string]func(client *shared.Client, userName string, chatID int64) error{
		"/hello":             sendHello,
		"/help":              sendHelp,
		"/commands":          sendCommands,
		"/track":             sendTrackHelp,
		"/report":            sendReportHelp,
		"/goal":              sendGoalHelp,
		"/wishlist":          sendWishlist,
		"/hatriki":           sendHatriki,
		"/tengohambre":       sendHambre,
		"/pinkipiensa":       sendPinki,
		"/chatID":            sendChatID,
		"/shoulddeploytoday": shoulddeploy.ShouldDeploy,
	}

	suffixReportMap = map[string]bool{
		"water":    true,
		"poop":     true,
		"keratine": true,
		"pipi":     true,
		"shower":   true,
		"run":      true,
		"tooth":    true,
		"sleep":    true,
		"read":     true,
		"gomita":   true,
		"all":      false,
		"monthly":  false,
	}

	suffixTrackMap = map[shared.Activity]bool{
		shared.Water:      true,
		shared.ToothBrush: true,
		shared.Read:       true,
		shared.Shower:     true,
		shared.Sleep:      true,
		shared.Gym:        true,
		shared.Poop:       true,
		shared.Run:        true,
		shared.Keratine:   true,
		shared.Pipi:       true,
		shared.Swimming:   true,
		shared.Cycling:    true,
		shared.Gomita:     true,
	}

	suffixGoalMap = map[string]func(client *shared.Client, userName, content string, chatID int64) error{
		"create": goals.SendCreateGoal,
		"delete": goals.SendDeleteGoal,
		"update": goals.SendUpdateGoal,
		"all":    goals.SendAllGoals,
	}

	prefixHandlers = map[string]func(ctx context.Context, client *shared.Client, chatID int64, userName, content string) error{
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
	-/track cycling
	-/track run
	-/track poop`

	msgHello = "Hola precioso \n\n" + msgHelp

	msgUnknownCommand = "q mondá es eso? 🤔"
)

// DoCommand handles the command
func DoCommand(ctx context.Context, client *shared.Client, chatID int64, userName string, command string) error {
	err := client.PrepareMenuButton(userName, chatID)
	if err != nil {
		return err
	}

	parts := strings.Split(command, " ")

	before, _, found := strings.Cut(parts[0], "@")
	if found {
		parts[0] = before
		command = strings.Join(parts, " ")
	}

	fn, ok := commandMap[command]
	if ok {
		return fn(client, userName, chatID)
	}

	for prefix, handler := range prefixHandlers {
		if strings.HasPrefix(command, prefix+" ") {
			content := strings.TrimPrefix(command, prefix+" ")

			return handler(ctx, client, chatID, userName, content)
		}
	}

	return sendUnknownCommand(client, chatID)
}

func handleTrack(ctx context.Context, client *shared.Client, chatID int64, userName, suffix string) error {
	before, after, _ := strings.Cut(suffix, " ")

	if _, ok := suffixTrackMap[shared.Activity(before)]; ok {
		return track.SendTrackActivity(ctx, shared.Activity(before), client, userName, after, chatID)
	}

	return sendUnknownCommand(client, chatID)
}

func handleReport(ctx context.Context, client *shared.Client, chatID int64, userName, suffix string) error {
	before, after, _ := strings.Cut(suffix, " ")

	if _, ok := suffixReportMap[before]; ok {
		return report.SendReportActicvity(ctx, client, shared.Activity(before), userName, after, chatID)
	}

	return sendUnknownCommand(client, chatID)
}

func handleGoal(ctx context.Context, client *shared.Client, chatID int64, userName, suffix string) error {
	before, after, _ := strings.Cut(suffix, " ")

	if fn, ok := suffixGoalMap[before]; ok {
		return fn(client, userName, after, chatID)
	}

	return sendUnknownCommand(client, chatID)
}

func sendUnknownCommand(client *shared.Client, chatID int64) error {
	return client.SendMessage(chatID, msgUnknownCommand)
}

func sendHello(client *shared.Client, userName string, chatID int64) error {
	return client.SendMessage(chatID, msgHello)
}

func sendHelp(client *shared.Client, userName string, chatID int64) error {
	return client.SendMessage(chatID, msgHelp)
}

func sendGoalHelp(client *shared.Client, userName string, chatID int64) error {
	return client.SendMessage(chatID, goals.MsgGoal)
}

func sendWishlist(client *shared.Client, userName string, chatID int64) error {
	ctx := context.Background()
	return wishlist.GetWishlist(ctx, client, userName, chatID)
}

func sendCommands(client *shared.Client, userName string, chatID int64) error {
	return client.SendMessage(chatID, msgCommands)
}

func sendTrackHelp(client *shared.Client, userName string, chatID int64) error {
	return client.SendMessage(chatID, msgTrack)
}

func sendReportHelp(client *shared.Client, username string, chatID int64) error {
	return client.SendMessage(chatID, "reporthelp")
}

func sendHatriki(client *shared.Client, userName string, chatID int64) error {
	return client.SendPhoto(chatID, hatriki)
}

func sendHambre(client *shared.Client, userName string, chatID int64) error {
	return client.SendPhoto(chatID, jeje)
}

func sendPinki(client *shared.Client, userName string, chatID int64) error {
	err := client.SendPhoto(chatID, pinki)
	if err != nil {
		return client.SendMessage(chatID, "vamos pinki piensa")
	}

	return nil
}

func sendChatID(client *shared.Client, userName string, chatID int64) error {
	message := fmt.Sprintf("Chat ID: %d", chatID)
	return client.SendMessage(chatID, message)
}
