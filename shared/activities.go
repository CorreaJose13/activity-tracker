package shared

import (
	"fmt"
	"time"
)

type Activity string

const (
	Water      Activity = "water"
	Gym        Activity = "gym"
	ToothBrush Activity = "toothbrush"
	Food       Activity = "food"
	Sleep      Activity = "sleep"
	Shower     Activity = "shower"
	Read       Activity = "read"
	Run        Activity = "run"
	Cycling    Activity = "cycling"
	Poop       Activity = "poop"
	Keratine   Activity = "keratine"
	Pipi       Activity = "pipi"
	Wishlist   Activity = "wishlist"
)

type Exercise string

const (
	Leg    Exercise = "leg"
	Bicep  Exercise = "bicep"
	Back   Exercise = "back"
	Tricep Exercise = "tricep"
	Abs    Exercise = "abs"
	Cardio Exercise = "cardio"
	Chest  Exercise = "chest"
)

// GetNow returns the current time in Colombia
func GetNow() (time.Time, error) {
	colombiaLocation, err := time.LoadLocation("America/Bogota")
	if err != nil {
		return time.Now(), fmt.Errorf(ErrGetNow, err.Error())
	}

	return time.Now().In(colombiaLocation), nil
}
