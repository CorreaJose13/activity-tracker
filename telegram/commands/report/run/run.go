package run

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"strconv"
	"time"
)

var (
	reportRunMessage = `Llegaron tus métricas de mierda subatleta %s 🏃🏾‍♂️

	Lunes: %s
	Martes: %s
	Miércoles: %s
	Jueves: %s
	Viernes: %s
	Sábado: %s
	Domingo: %s

	Si no querés que te robe alguien de sucre comenzá a correr más 🤢
	`

	runninGif = "https://media.giphy.com/media/XyNMDllviurb3bpfDE/giphy.gif?cid=ecf05e47xnoxtwqok8a12t2uqy5bsqr1z0iwo057vl78ik3b&ep=v1_gifs_search&rid=giphy.gif&ct=g"
)

func getRunString(count float64) string {
	if count == 0 {
		return "ni un metro, perezoso 🤢"
	}

	return fmt.Sprintf("%.2f", count) + " kms 🏃🏾‍♂️"
}

// SendRunReport sends the run report
func SendRunReport(client *shared.Client, userName, content string, chatID int64) error {
	report, err := GenerateRunReport(client, userName, chatID)
	if err != nil {
		return err
	}

	err = client.SendMessage(chatID, report)
	if err != nil {
		return err
	}

	return client.SendAnimation(chatID, runninGif)
}

// GenerateRunReport generates the run report
func GenerateRunReport(client *shared.Client, userName string, chatID int64) (string, error) {
	runActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, shared.Run)
	if err != nil {
		return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	runAmount := map[time.Weekday]float64{
		time.Monday:    0,
		time.Tuesday:   0,
		time.Wednesday: 0,
		time.Thursday:  0,
		time.Friday:    0,
		time.Saturday:  0,
		time.Sunday:    0,
	}

	for _, activity := range runActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		num, err := strconv.ParseFloat(activity.Content, 64)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		runAmount[day] += num
	}

	report := fmt.Sprintf(reportRunMessage, userName,
		getRunString(runAmount[time.Monday]),
		getRunString(runAmount[time.Tuesday]),
		getRunString(runAmount[time.Wednesday]),
		getRunString(runAmount[time.Thursday]),
		getRunString(runAmount[time.Friday]),
		getRunString(runAmount[time.Saturday]),
		getRunString(runAmount[time.Sunday]))

	return report, nil
}
