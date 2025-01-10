package shared

import (
	"fmt"
	"time"
)

type Activity string

type Exercise string

// UserActivity contains an user activity info
type UserActivity struct {
	ID           string   `bson:"id"`
	Name         string   `bson:"name"`
	Activity     Activity `bson:"activity"`
	ExerciseType Exercise `bson:"excercise_type,omitempty"`
	CreatedAt    string   `bson:"created_at"`
	UpdatedAt    string   `bson:"updated_at"`
	Content      string   `bson:"content,omitempty"`
}

const (
	Water      Activity = "water"
	Gym        Activity = "gym"
	ToothBrush Activity = "toothbrush"
	Food       Activity = "food"
	Sleep      Activity = "sleep"
	Shower     Activity = "shower"
	Read       Activity = "read"
	Run        Activity = "run"
	Swimming   Activity = "swimming"
	Cycling    Activity = "cycling"
	Poop       Activity = "poop"
	Keratine   Activity = "keratine"
	Pipi       Activity = "pipi"
	Wishlist   Activity = "wishlist"
)

const (
	Leg    Exercise = "leg"
	Bicep  Exercise = "bicep"
	Back   Exercise = "back"
	Tricep Exercise = "tricep"
	Abs    Exercise = "abs"
	Cardio Exercise = "cardio"
	Chest  Exercise = "chest"
)

// NewActivity creates a new activity with timestamp in format time.RFC3339
func NewActivity(activity Activity, userName string, content string) (UserActivity, error) {
	now, err := GetNow()
	if err != nil {
		return UserActivity{}, err
	}

	nowStr := now.Format(time.RFC3339)

	return UserActivity{
		ID:        GenerateActivityItemID(now, userName, activity),
		Name:      userName,
		Activity:  activity,
		Content:   content,
		CreatedAt: nowStr,
	}, nil
}

// GetNow returns the current time in Colombia
func GetNow() (time.Time, error) {
	colombiaLocation, err := time.LoadLocation("America/Bogota")
	if err != nil {
		return time.Time{}, fmt.Errorf(ErrGetNow, err.Error())
	}

	return time.Now().In(colombiaLocation), nil
}

// GenerateActivityItemID generate the unique id of the activity item that will be saved in the activity database
func GenerateActivityItemID(now time.Time, username string, activity Activity) string {
	formattedNow := now.Format(time.RFC3339)

	return fmt.Sprintf("%s-%s-%s", formattedNow, username, activity)
}
