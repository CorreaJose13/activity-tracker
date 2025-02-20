package shoulddeploy

import (
	"activity-tracker/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	shouldDeployURL   = "https://shouldideploy.today/api/slack?tz=Etc/GMT-5"
	finalMessage      = "Le entregamos tu jijueputa diploi a la bola mágica 🔮 y esto respondió:\n\n%s"
	requestFailedMsg  = "No je pudo hacer la request: %s"
	failedResponseMsg = "la bola mágica se mimió, status: %d"
	noResultsMsg      = "esta vez la bola je apagó :("
)

type ShouldDeployResponse struct {
	ResponseType string       `json:"response_type"`
	Attachments  []Attachment `json:"attachments"`
}

type Attachment struct {
	Text       string `json:"text"`
	Color      string `json:"color"`
	ThumbURL   string `json:"thumb_url"`
	FooterIcon string `json:"footer_icon"`
	Footer     string `json:"footer"`
}

// ShouldDeploy sends a message to the user with the result of the shouldideploy.today API
func ShouldDeploy(client *shared.Client, userName string, chatID int64) error {
	resp, err := http.Get(shouldDeployURL)
	if err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(requestFailedMsg, err.Error()))
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return client.SendMessage(chatID, fmt.Sprintf(failedResponseMsg, resp.StatusCode))
	}

	var deployResp ShouldDeployResponse
	if err := json.NewDecoder(resp.Body).Decode(&deployResp); err != nil {
		return client.SendMessage(chatID, fmt.Sprintf(requestFailedMsg, err.Error()))
	}

	if len(deployResp.Attachments) == 0 {
		return client.SendMessage(chatID, noResultsMsg)
	}

	message := deployResp.Attachments[0].Text
	if message == "" {
		return client.SendMessage(chatID, noResultsMsg)
	}

	return client.SendMessage(chatID, fmt.Sprintf(finalMessage, message))
}
