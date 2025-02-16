package gomita

import (
	"activity-tracker/shared"
	"activity-tracker/storage"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

var (
	reportMessage = `Aquí tienes un resumen de cuantas veces te trabaste con gomita esta semana  %s 🍬🌿
	
	De acuerdo a la cochinada de base de datos que tenemos, esto es lo que encontré:

	Lunes: %s 
	Martes: %s 
	Miércoles: %s 
	Jueves: %s 
	Viernes: %s 
	Sábado: %s 
	Domingo: %s 

	Ponte las pilas pa consumir más 🥵
	`

	gomitaGif = "https://media.giphy.com/media/rWiEbamfqOHrq/giphy.gif?cid=790b76113to9r3hgdf5j4317otxivd5ldi4gg7pug36eam97&ep=v1_gifs_search&rid=giphy.gif&ct=g"
)

func getGomitaString(count float64) string {
	if count == 0 {
		return "ni una bb :("
	}

	if count <= 1 {
		p, q := floatToFraction(count)
		return fmt.Sprintf("%d/%d", p, q) + " de gomita 🍁"
	}

	if count == 1 {
		return "1 gomita 🍁"
	}

	return fmt.Sprintf("%.2f", count) + " gomitas 🍁"
}

func floatToFraction(f float64) (p, q int64) {
	// Create a new big.Float from the float64 value
	bf := big.NewFloat(f)

	// Set the precision high enough to handle the float
	bf.SetPrec(64)

	// Convert the big.Float to a big.Rat (rational number)
	rat := new(big.Rat)
	rat.SetFloat64(f)

	// Get the numerator and denominator
	p = rat.Num().Int64()
	q = rat.Denom().Int64()

	return p, q
}

func SendGomitaReport(client *shared.Client, userName, content string, chatID int64) error {
	report, err := GenerateGomitaReport(client, userName, chatID)
	if err != nil {
		return err
	}

	err = client.SendMessage(chatID, report)
	if err != nil {
		return err
	}

	return client.SendAnimation(chatID, gomitaGif)
}

func GenerateGomitaReport(client *shared.Client, userName string, chatID int64) (string, error) {
	gomitaActivities, err := storage.GetLastWeekUserHistoryPerActivity(userName, shared.Gomita)
	if err != nil {
		return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
	}

	gomitasPerDay := map[time.Weekday]float64{
		time.Monday:    0,
		time.Tuesday:   0,
		time.Wednesday: 0,
		time.Thursday:  0,
		time.Friday:    0,
		time.Saturday:  0,
		time.Sunday:    0,
	}

	for _, activity := range gomitaActivities {
		createdAt, err := time.Parse(time.RFC3339, activity.CreatedAt)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		day := createdAt.Weekday()

		num, err := strconv.ParseFloat(activity.Content, 64)
		if err != nil {
			return "", client.SendMessage(chatID, fmt.Sprintf(shared.ErrSendMessage, err.Error()))
		}

		gomitasPerDay[day] += num
	}

	report := fmt.Sprintf(reportMessage, userName,
		getGomitaString(gomitasPerDay[time.Monday]),
		getGomitaString(gomitasPerDay[time.Tuesday]),
		getGomitaString(gomitasPerDay[time.Wednesday]),
		getGomitaString(gomitasPerDay[time.Thursday]),
		getGomitaString(gomitasPerDay[time.Friday]),
		getGomitaString(gomitasPerDay[time.Saturday]),
		getGomitaString(gomitasPerDay[time.Sunday]))

	return report, nil
}
