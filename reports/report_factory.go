package reports

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"
)

var (
	ErrReportNotImplemented = errors.New("report not implemented")
	ErrEmptyActivities      = errors.New("empty activities")
	mapReportByActivity     = map[shared.Activity]func(string) (Report, error){
		shared.Gomita:     NewGomitaReport,
		shared.Keratine:   NewKeratineReport,
		shared.Pipi:       NewPipiReport,
		shared.Run:        NewRunReport,
		shared.Shower:     NewShowerReport,
		shared.Sleep:      NewSleepReport,
		shared.ToothBrush: NewToothReport,
		shared.Water:      NewWaterReport,
	}

	getLastWeekUserHistoryPerActivity = storage.GetLastWeekUserHistoryPerActivity
)

type DailyReport struct {
	Date   time.Time
	Result float64
}

type SourceType string
type Period string

const (
	APISource SourceType = "api"
	TGSource  SourceType = "telegram"

	PeriodWeekly = "weekly"
)

type ReportError struct {
	BaseError error
	Details   string
}

func (e *ReportError) Error() string {
	return fmt.Sprintf("%s: %s", e.BaseError.Error(), e.Details)
}

func (e *ReportError) Unwrap() error {
	return e.BaseError
}

type Report interface {
	SetDailyReports([]DailyReport)
	String() string
	DailyReports() []DailyReport
	GetGifURL() string
	GetPeriod() Period
	GetActivity() shared.Activity
	GetUsername() string
}

func NewReport(activityType shared.Activity, username string) (Report, error) {
	newReport, ok := mapReportByActivity[activityType]
	if !ok {
		return nil, &ReportError{
			BaseError: ErrReportNotImplemented,
			Details:   fmt.Sprintf("[%s]", activityType),
		}
	}

	return newReport(username)
}

func PanicReportNotGenerated() {
	panic("reports were not generated")
}

func AddDailyReports(ctx context.Context, report Report) error {
	activities, err := activitiesByPeriod(ctx, report)
	if err != nil {
		return err
	}

	dailyReports, err := activitiesToDailyReports(activities)
	if err != nil {
		return err
	}

	report.SetDailyReports(dailyReports)

	return nil
}

func activitiesToDailyReports(activities []*shared.UserActivity) ([]DailyReport, error) {
	resultPerDate := map[time.Time]DailyReport{}

	for _, activity := range activities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return []DailyReport{}, err
		}

		num, err := strconv.ParseFloat(activity.Content, 64)
		if err != nil {
			return []DailyReport{}, err
		}

		newDate := time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), 0, 0, 0, 0, createdAt.Location())
		resultPerDate[newDate] = DailyReport{
			Date:   newDate,
			Result: resultPerDate[newDate].Result + num,
		}
	}

	result := []DailyReport{}

	for _, dailyReport := range resultPerDate {
		result = append(result, dailyReport)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Date.Before(result[j].Date)
	})

	return result, nil
}

func activitiesByPeriod(ctx context.Context, report Report) ([]*shared.UserActivity, error) {
	err := ErrEmptyActivities
	var activities []*shared.UserActivity

	if report.GetPeriod() == PeriodWeekly {
		activities, err = getLastWeekUserHistoryPerActivity(ctx, report.GetUsername(), report.GetActivity())
	}

	return activities, err
}
