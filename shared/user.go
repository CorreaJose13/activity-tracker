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

	AdminUsersChatIDs = map[string]int64{
		Brayan: BrayanChatID,
	}
)

// User contains an user info
type User struct {
	Name             string     `bson:"name"`
	ChatID           int64      `bson:"chat_id"`
	EnabledActivites []Activity `bson:"enabled_activities"`
}

// NewUser creates a new user
func NewUser(name string, chatID int64, enabledActivities []Activity) *User {
	return &User{
		Name:             name,
		ChatID:           chatID,
		EnabledActivites: enabledActivities,
	}
}

// GetRandomUserName returns a random user name from the list
func GetRandomUserName() string {
	users := []string{Valery, Brayan, Johan, Jose, Daniela, Juan}

	return users[rand.Intn(len(users))]
}
