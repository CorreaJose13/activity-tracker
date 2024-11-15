package telegram

import (
	"activity-tracker/api/telegram"
	"activity-tracker/shared"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendTrackPipi(t *testing.T) {
	c := require.New(t)

	err := sendTrackPipi(&telegram.Bot{}, "test", "", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}

func TestSendTrackKeratine(t *testing.T) {
	c := require.New(t)

	err := sendTrackKeratine(&telegram.Bot{}, "test1", "", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}

func TestSendTrackWater(t *testing.T) {
	c := require.New(t)

	err := sendTrackWater(&telegram.Bot{}, "test2", "", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}

func TestSendTrackTooth(t *testing.T) {
	c := require.New(t)

	err := sendTrackTooth(&telegram.Bot{}, "test3", "", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}

func TestSendTrackRun(t *testing.T) {
	c := require.New(t)

	err := sendTrackRun(&telegram.Bot{}, "test4", "5km", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}

func TestSendCreateGoal(t *testing.T) {
	c := require.New(t)

	err := sendCreateGoal(&telegram.Bot{}, "test", "", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}

func TestSendDeleteGoal(t *testing.T) {
	c := require.New(t)

	err := sendDeleteGoal(&telegram.Bot{}, "test", "", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}

func TestSendUpdateGoal(t *testing.T) {
	c := require.New(t)

	err := sendUpdateGoal(&telegram.Bot{}, "test", "", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}

func TestSendAllGoals(t *testing.T) {
	c := require.New(t)

	err := sendAllGoals(&telegram.Bot{}, "test", "", 1)

	// It's not necessary to test the send message error
	c.Error(err)

	// It is set to fail to check the trace code in commands.go
	c.Equal("a", "b")
}

func TestParseGoalsToString(t *testing.T) {
	c := require.New(t)

	goals := []*shared.PersonalGoal{
		{Activity: "test", GoalConfig: shared.GoalConfig{
			Daily:   "test",
			Weekly:  "test1",
			Monthly: "test2",
		}},
	}

	result := parseGoalsToString(&telegram.Bot{}, 1, goals)
	c.Contains(result, "test")
	c.Contains(result, "test1")
	c.Contains(result, "test2")
}
