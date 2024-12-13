package goals

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"encoding/json"
	"fmt"
	"strings"
)

var (
	msgGoal = `por fin vas a querer hacer algo diferente 🐸🐶
haceme el cruce y me decís qué querés con un objetivo

algo así precioso:

crear un objetivo: /goal create <activity> <goal> <optional: --daily=VALUE> <optional: --weekly=VALUE> <optional: --monthly=VALUE>	
borrar un objetivo: /goal delete <activity>
actualizar un objetivo: /goal update <activity> <new goal>
obtener todos tus objetivos: /goal all
`

	msgAllGoals = `tus objetivos son:\n\n%s`
)

// SendAllGoals handles the all goals message
func SendAllGoals(client *shared.Client, userName, content string, chatID int64) error {
	goals, err := storage.GetAllPersonalGoals(userName)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	if len(goals) == 0 {
		return client.SendMessage(chatID, fmt.Sprintf("carecemos de objetivos bb, animate a crear uno y redimes un besucio de @%s", shared.GetRandomUserName()))
	}

	msg := parseGoalsToString(client, chatID, goals)

	return client.SendMessage(chatID, fmt.Sprintf(msgAllGoals, msg))
}

func parseGoalsToString(client *shared.Client, chatID int64, goals []*shared.PersonalGoal) string {
	var result strings.Builder
	for _, goal := range goals {
		goalConfigJSON, err := json.Marshal(goal.GoalConfig)
		if err != nil {
			_ = client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))

			continue
		}

		result.WriteString(fmt.Sprintf("- %s: %s\n", goal.Activity, string(goalConfigJSON)))
	}

	return result.String()
}
