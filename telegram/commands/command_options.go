package commands

import (
	"activity-tracker/shared"
	"activity-tracker/telegram/commands/gemini"
	goals "activity-tracker/telegram/commands/goals"
	shoulddeploy "activity-tracker/telegram/commands/shoulddeploy"
	"activity-tracker/telegram/commands/wishlist"
	"context"
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
)

func GetAvailableActivitiesToReport() []shared.Activity {
	activities := []shared.Activity{}

	for activity, available := range suffixReportMap {
		if available {
			activities = append(activities, shared.Activity(activity))
		}
	}

	return activities
}
