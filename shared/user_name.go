package shared

import (
	"golang.org/x/exp/rand"
)

const (
	Valery = "ValeryMolinaB"
	Brayan = "BrayanEscobar"
	Johan  = "JohanFlorez"
	Jose   = "jCorreaM"
	Mauro  = "mcortazar"
)

// GetRandomUserName returns a random user name from the list
func GetRandomUserName() string {
	users := []string{Valery, Brayan, Johan, Jose, Mauro}

	return users[rand.Intn(len(users))]
}
