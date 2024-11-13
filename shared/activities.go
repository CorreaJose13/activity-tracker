package shared

import "time"

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
func GetNow() time.Time {
	colombiaLocation, err := time.LoadLocation("America/Bogota")
	if err != nil {
		panic("failed to load Colombia timezone: " + err.Error())
	}

	return time.Now().In(colombiaLocation)
}
