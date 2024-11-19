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

algo así tetranutra:

crear un objetivo: /goal create <activity> <goal> <optional: --daily=VALUE> <optional: --weekly=VALUE> <optional: --monthly=VALUE>	
borrar un objetivo: /goal delete <activity>
actualizar un objetivo: /goal update <activity> <new goal>
obtener todos tus objetivos: /goal all
`

	msgAllGoals = `tus objetivos son:\n\n%s`
)

// SendAllGoals handles the all goals message
func SendAllGoals(bot *shared.Bot, userName, content string, chatID int64) error {
	goals, err := storage.GetAllPersonalGoals(userName)
	if err != nil {
		return shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	if len(goals) == 0 {
		return shared.SendMessage(bot, chatID, fmt.Sprintf("carecemos de objetivos bb, animate a crear uno y redimes un besucio de @%s", shared.GetRandomUserName()))
	}

	msg := parseGoalsToString(bot, chatID, goals)

	return shared.SendMessage(bot, chatID, fmt.Sprintf(msgAllGoals, msg))
}

func parseGoalsToString(bot *shared.Bot, chatID int64, goals []*shared.PersonalGoal) string {
	var result strings.Builder
	for _, goal := range goals {
		goalConfigJSON, err := json.Marshal(goal.GoalConfig)
		if err != nil {
			_ = shared.SendMessage(bot, chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))

			continue
		}

		result.WriteString(fmt.Sprintf("- %s: %s\n", goal.Activity, string(goalConfigJSON)))
	}

	return result.String()
}
