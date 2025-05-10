package reports

import (
	"activity-tracker/shared"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewReportSuccess(t *testing.T) {
	c := require.New(t)

	_, err := NewReport(shared.Gomita, "dummy")
	c.NoError(err)
}

func TestNewReportFailed(t *testing.T) {
	c := require.New(t)

	_, err := NewReport("dummy", "dummy")
	c.ErrorIs(err, ErrReportNotImplemented)
}

func TestActivityToReportsSuccess(t *testing.T) {
	c := require.New(t)

	now := time.Now().Format(time.RFC3339)

	activities := []*shared.UserActivity{
		{
			CreatedAt: now,
			Activity:  shared.Gomita,
			Content:   "2",
		},
		{
			CreatedAt: now,
			Activity:  shared.Gomita,
			Content:   "2",
		},
	}

	dailyReports, err := activitiesToDailyReports(activities)
	c.NoError(err)

	c.Len(dailyReports, 1)
	c.Equal(dailyReports[0].Result, float64(4))
}

func TestAddDailyReportsSuccess(t *testing.T) {
	c := require.New(t)

	now := time.Now().Format(time.RFC3339)

	activities := []*shared.UserActivity{
		{
			CreatedAt: now,
			Activity:  shared.Gomita,
			Content:   "2",
		},
		{
			CreatedAt: now,
			Activity:  shared.Gomita,
			Content:   "2",
		},
	}

	getLastWeekUserHistoryPerActivity = func(_ context.Context, _ string, _ shared.Activity) ([]*shared.UserActivity, error) {
		return activities, nil
	}

	rp, err := NewReport(shared.Gomita, "dummy")
	c.NoError(err)

	err = AddDailyReports(context.Background(), rp)
	c.NoError(err)

	c.Len(rp.DailyReports(), 1)
}

func TestGenerateReports(t *testing.T) {
	c := require.New(t)

	now := time.Now().Format(time.RFC3339)

	activities := []*shared.UserActivity{
		{
			CreatedAt: now,
			Activity:  shared.Gomita,
			Content:   "2",
		},
		{
			CreatedAt: now,
			Activity:  shared.Keratine,
			Content:   "1",
		},
		{
			CreatedAt: now,
			Activity:  shared.Pipi,
			Content:   "1",
		},
		{
			CreatedAt: now,
			Activity:  shared.Run,
			Content:   "5",
		},
		{
			CreatedAt: now,
			Activity:  shared.Shower,
			Content:   "1",
		},
		{
			CreatedAt: now,
			Activity:  shared.Sleep,
			Content:   "8",
		},
		{
			CreatedAt: now,
			Activity:  shared.ToothBrush,
			Content:   "2",
		},
		{
			CreatedAt: now,
			Activity:  shared.Water,
			Content:   "3",
		},
	}

	getLastWeekUserHistoryPerActivity = func(_ context.Context, _ string, _ shared.Activity) ([]*shared.UserActivity, error) {
		return activities, nil
	}

	for _, ac := range activities {
		rp, err := NewReport(ac.Activity, "dummy")
		c.NoError(err)

		err = AddDailyReports(context.Background(), rp)
		c.NoError(err)

		c.NotEmpty(rp.String())
		c.Len(rp.DailyReports(), 1)

		if rp.GetGifURL() == "" {
			t.Error(fmt.Sprintf("%s gif not should be empty", ac.Activity))
		}
	}
}
