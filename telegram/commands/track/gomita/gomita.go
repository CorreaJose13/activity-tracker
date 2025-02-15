package gomita

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"strconv"
	"strings"
)

var (
	missingContentMessage = "tan trabao estás?? mandame la info así vicioso:\n\t /track gomita <fraction or float>"
	invalidContentMessage = "concentrate hijo de toda tu puta madre, escribe bien el número 😡"
	successMessage        = "buen biaje mi sielo 🌈"
)

func fractionToFloatString(content string) (string, error) {
	split := strings.Split(content, "/")
	num, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		return "", err
	}
	den, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
		return "", err
	}
	result := num / den
	floatString := fmt.Sprintf("%f", result)
	return floatString, nil
}

// SendTrackGomita tracks how many gomitas did u take
func SendTrackGomita(client *shared.Client, userName, content string, chatID int64) error {
	if content == "" {
		return client.SendMessage(chatID, missingContentMessage)
	}

	if strings.Contains(content, "/") {
		floatString, err := fractionToFloatString(content)
		if err != nil {
		}
		content = floatString
	} else if !shared.IsValidFloat(content) {
		return client.SendMessage(chatID, invalidContentMessage)
	}

	userActivity, err := shared.NewActivity(shared.Gomita, userName, content)
	if err != nil {
		return client.SendMessage(chatID, err.Error())
	}

	err = storage.Create(userActivity)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	return client.SendMessage(chatID, successMessage)
}
