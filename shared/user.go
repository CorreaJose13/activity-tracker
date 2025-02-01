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

	// TODO: Add the missing user chat ids in another PR
	ValeryChatID = int64(2071323849)
	BrayanChatID = int64(622185634)
	JohanChatID  = int64(815505945)
	JoseChatID   = int64(2112190700)

	// Schedulers chat ids
	KeratineSchedulerChatIDs = map[string]int64{
		Valery: ValeryChatID,
		Brayan: BrayanChatID,
		Johan:  JohanChatID,
	}

	AllReportsSchedulerChatIDs = map[string]int64{
		Valery: ValeryChatID,
		Brayan: BrayanChatID,
		Johan:  JohanChatID,
		Jose:   JoseChatID,
	}
)

// GetRandomUserName returns a random user name from the list
func GetRandomUserName() string {
	users := []string{Valery, Brayan, Johan, Jose, Daniela, Juan}

	return users[rand.Intn(len(users))]
}
