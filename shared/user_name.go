package shared

import (
	"golang.org/x/exp/rand"
)

var (
	// User names
	Valery  = "ValeryMolinaB"
	Brayan  = "BrayanEscobar"
	Johan   = "JohanFlorez"
	Jose    = "jCorreaM"
	Juan    = "juancballesteros"
	Daniela = "DaniCD8"

	ValeryChatID = int64(2071323849)
	BrayanChatID = int64(622185634)
	JohanChatID  = int64(815505945)

	// User chat IDs
	UsersChatIDs = map[string]int64{
		// Valery:  ValeryChatID,
		Brayan: BrayanChatID,
		// Johan:   JohanChatID,
	}

	// Keratine chat IDs
	KeratineChatIDs = []int64{ValeryChatID, BrayanChatID, JohanChatID}

	// All reports chat IDs
	AllReportsChatIDs = []int64{ValeryChatID, BrayanChatID, JohanChatID}
)

// GetRandomUserName returns a random user name from the list
func GetRandomUserName() string {
	users := []string{Valery, Brayan, Johan, Jose, Daniela, Juan}

	return users[rand.Intn(len(users))]
}
